package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"strconv"
	"time"

	binance_struct "github.com/IEatLemons/goUtils/trading/binance/struct"
)

type klines struct {
	klines *binance
}

func InitKLines(env Environment) *klines {
	return &klines{
		klines: NewBinance(env, "/api/v3/klines"),
	}
}

type KlinesReq struct {
	Symbol    Symbol    `json:"symbol"`
	Interval  Interval  `json:"interval"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Limit     uint64    `json:"limit"`
}

func (c *klines) GetKlines(params *KlinesReq) (Klines []*binance_struct.Klines, err error) {
	values := &url.Values{}
	values.Set("symbol", string(params.Symbol))
	values.Set("interval", string(params.Interval))
	if params.Limit != 0 {
		values.Set("limit", strconv.FormatUint(params.Limit, 10))
	}
	result, err := c.klines.QuestSimple(values)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(result))
	var data [][]interface{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("Parsing JSON failed:", err)
		return
	}
	Klines = []*binance_struct.Klines{}
	for _, data := range data {
		OpenPrice, _ := strconv.ParseFloat(data[1].(string), 64)
		HighPrice, _ := strconv.ParseFloat(data[2].(string), 64)
		LowPrice, _ := strconv.ParseFloat(data[3].(string), 64)
		ClosePrice, _ := strconv.ParseFloat(data[4].(string), 64)
		Volume, _ := strconv.ParseFloat(data[5].(string), 64)
		VolumeOfBusiness, _ := strconv.ParseFloat(data[7].(string), 64)
		ActiveBuyingVolume, _ := strconv.ParseFloat(data[9].(string), 64)
		ActiveBuyingTurnover, _ := strconv.ParseFloat(data[10].(string), 64)
		Ignore, _ := strconv.ParseFloat(data[11].(string), 64)
		Kline := &binance_struct.Klines{
			OpenTime:             time.Unix(int64((data[0].(float64) / 1000)), 0),
			OpenPrice:            big.NewFloat(OpenPrice),
			HighPrice:            big.NewFloat(HighPrice),
			LowPrice:             big.NewFloat(LowPrice),
			ClosePrice:           big.NewFloat(ClosePrice),
			Volume:               big.NewFloat(Volume),
			ClosingTime:          time.Unix(int64((data[6].(float64) / 1000)), 0),
			VolumeOfBusiness:     big.NewFloat(VolumeOfBusiness),
			TxNumber:             int(data[8].(float64)),
			ActiveBuyingVolume:   big.NewFloat(ActiveBuyingVolume),
			ActiveBuyingTurnover: big.NewFloat(ActiveBuyingTurnover),
			Ignore:               big.NewFloat(Ignore),
		}

		Klines = append(Klines, Kline)
	}

	return
}
