package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/IEatLemons/goUtils/request"
)

const (
	KlinesPath PATH = "/api/v3/klines"
)

type KlinesReq struct {
	Symbol   Symbol   `json:"symbol"`
	Interval Interval `json:"interval"`
}

type KlinesData struct {
	Time       time.Time `json:"0"`
	OpenPrice  float64   `json:"1"`
	HighPrice  float64   `json:"2"`
	LowPrice   float64   `json:"3"`
	ClosePrice float64   `json:"4"`
	Volume     float64   `json:"5"`
}

func GetKlines(params *KlinesReq) (Klines []*KlinesData) {
	reqPar, err := request.Struct2Params(params)
	if err != nil {
		log.Fatalln(err)
	}
	result, err := request.NewRequest(APIURL+string(KlinesPath), []request.RequestOptions{
		request.SetParams(reqPar),
	}...).Send()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(result))
	var data [][]interface{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("Parsing JSON failed:", err)
		return
	}
	Klines = []*KlinesData{}
	for _, data := range data {
		openPrice := data[1].(string) // At the opening
		OpenPrice, _ := strconv.ParseFloat(openPrice, 64)
		highPrice := data[2].(string) // The highest price
		HighPrice, _ := strconv.ParseFloat(highPrice, 64)
		lowPrice := data[3].(string) // Bottom price
		LowPrice, _ := strconv.ParseFloat(lowPrice, 64)
		closePrice := data[4].(string) // Closing price
		ClosePrice, _ := strconv.ParseFloat(closePrice, 64)
		volume := data[5].(string) // Volume of transaction
		Volume, _ := strconv.ParseFloat(volume, 64)
		Kline := &KlinesData{
			Time:       time.Unix(int64((data[0].(float64) / 1000)), 0),
			OpenPrice:  OpenPrice,
			HighPrice:  HighPrice,
			LowPrice:   LowPrice,
			ClosePrice: ClosePrice,
			Volume:     Volume,
		}

		Klines = append(Klines, Kline)
	}

	return
}


