package binance

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type Stream struct {
	Method Method
	Params MonitorParam
}

type MonitorParam interface {
	Symbol() Symbol
	Value() string
}

type Processing interface {
	Successful([]byte)
}

type monitor struct {
	spotWs *binance
	conn   *websocket.Conn
}

func InitMonitor(env Environment) (*monitor, error) {
	spotWs := NewBinance(env, "")
	conn, _, err := websocket.DefaultDialer.Dial(string(spotWs.SpotWss), nil)
	if err != nil {
		return nil, err
	}
	m := &monitor{
		spotWs: NewBinance(env, ""),
		conn:   conn,
	}
	// defer m.conn.Close()
	return m, nil
}

func (c *monitor) Close() {
	c.conn.Close()
}

func (c *monitor) Monitor(Proc Processing, Streams ...Stream) {
	if len(Streams) < 1 {
		log.Fatal("[Binance]The listener must be larger than 1")
	}

	for _, Stream := range Streams {
		subscribeMsg := getSubscribeMsg(Stream)
		// subscribeMsg := fmt.Sprintf(`{"method": "SUBSCRIBE", "params": ["%s@kline_1m"], "id": 1}`, "btcusdt")
		if err := c.conn.WriteMessage(websocket.TextMessage, []byte(subscribeMsg)); err != nil {
			log.Fatal("[Binance]Sending subscription message error:", err)
		}
	}

	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("[Binance]Read message error:", err)
				return
			}
			Proc.Successful(message)
		}
	}()
}

func getSubscribeMsg(Stream Stream) string {
	return fmt.Sprintf(`{"method": "%s", "params": ["%s@%s"], "id": 1}`, Stream.Method, strings.ToLower(string(Stream.Params.Symbol())), Stream.Params.Value())
}

type KLineMonitor struct {
	S Symbol
	I Interval
}

func (M KLineMonitor) Symbol() Symbol {
	return M.S
}
func (M KLineMonitor) Value() string {
	return "kline_" + string(M.I)
}

type MiniTickerMonitor struct {
	S Symbol
}

func (M MiniTickerMonitor) Symbol() Symbol {
	return M.S
}
func (M MiniTickerMonitor) Value() string {
	return "miniTicker"
}
