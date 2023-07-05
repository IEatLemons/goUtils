package binance_struct

import (
	"math/big"
	"time"
)

type Klines struct {
	OpenTime             time.Time  `json:"0"`  // k线开盘时间
	OpenPrice            *big.Float `json:"1"`  // 开盘价
	HighPrice            *big.Float `json:"2"`  // 最高价
	LowPrice             *big.Float `json:"3"`  // 最低价
	ClosePrice           *big.Float `json:"4"`  // 收盘价
	Volume               *big.Float `json:"5"`  // 成交量
	ClosingTime          time.Time  `json:"6"`  // 收盘价(当前K线未结束的即为最新价)
	VolumeOfBusiness     *big.Float `json:"7"`  // 成交额
	TxNumber             int        `json:"8"`  // 成交笔数
	ActiveBuyingVolume   *big.Float `json:"9"`  // 主动买入成交量
	ActiveBuyingTurnover *big.Float `json:"10"` // 主动买入成交额
	Ignore               *big.Float `json:"11"` // 请忽略该参数
}
