package binance_struct

import (
	"math/big"
	"time"
)

type ExchangeInfo struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []*RateLimit  `json:"rateLimits"` // API访问的限制
	ServerTime      time.Time     `json:"serverTime"` // 请忽略。如果需要获取当前系统时间，请查询接口 "GET /fapi/v1/time"
	Assets          []*Asset      `json:"assets"`     // 资产信息
	Symbols         []*Symbol     `json:"symbols"`    // 交易对信息
	Timezone        string        `json:"timezone"`   // 服务器所用的时间区域
}

type RateLimit struct {
	Interval      string `json:"interval"`      // 按照分钟计算
	IntervalNum   int    `json:"intervalNum"`   // 按照1分钟计算
	Limit         int    `json:"limit"`         // 上限次数
	RateLimitType string `json:"rateLimitType"` // 按照访问权重来计算
}
type Asset struct {
	Asset             string `json:"asset"`             //  资产名称
	MarginAvailable   bool   `json:"marginAvailable"`   // 是否可用作保证金
	AutoAssetExchange int    `json:"autoAssetExchange"` // 保证金资产自动兑换阈值
}
type Symbol struct {
	Symbol                string     `json:"symbol"`                // 交易对
	Pair                  string     `json:"pair"`                  // 标的交易对
	ContractType          string     `json:"contractType"`          // 合约类型
	DeliveryDate          time.Time  `json:"deliveryDate"`          // 交割日期
	OnboardDate           time.Time  `json:"onboardDate"`           // 上线日期
	Status                string     `json:"status"`                // 交易对状态
	MaintMarginPercent    *big.Float `json:"maintMarginPercent"`    // 请忽略
	RequiredMarginPercent *big.Float `json:"requiredMarginPercent"` // 请忽略
	BaseAsset             string     `json:"baseAsset"`             // 标的资产
	QuoteAsset            string     `json:"quoteAsset"`            // 报价资产
	MarginAsset           string     `json:"marginAsset"`           // 保证金资产
	PricePrecision        int        `json:"pricePrecision"`        // 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分）
	QuantityPrecision     int        `json:"quantityPrecision"`     // 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分）
	BaseAssetPrecision    int        `json:"baseAssetPrecision"`    // 标的资产精度
	QuotePrecision        int        `json:"quotePrecision"`        // 报价资产精度
	UnderlyingType        string     `json:"underlyingType"`
	UnderlyingSubType     []string   `json:"underlyingSubType"`
	SettlePlan            int        `json:"settlePlan"`
	TriggerProtect        *big.Float `json:"triggerProtect"` // 开启"priceProtect"的条件订单的触发阈值
	Filters               []*filter  `json:"filters"`

	//	"OrderType": [ // 订单类型
	// 		"LIMIT",  // 限价单
	// 		"MARKET",  // 市价单
	//		"STOP", // 止损单
	// 		"STOP_MARKET", // 止损市价单
	//		"TAKE_PROFIT", // 止盈单
	// 		"TAKE_PROFIT_MARKET", // 止盈暑市价单
	//     	"TRAILING_STOP_MARKET" // 跟踪止损市价单
	//	]
	OrderType []string `json:"OrderType"` // 订单类型

	//	"timeInForce": [ // 有效方式
	//  	"GTC", // 成交为止, 一直有效
	//     	"IOC", // 无法立即成交(吃单)的部分就撤销
	//    	"FOK", // 无法全部立即成交就撤销
	//  	"GTX" // 无法成为挂单方就撤销
	//	],
	TimeInForce     []string   `json:"timeInForce"`     // 订单类型
	LiquidationFee  *big.Float `json:"liquidationFee"`  // 强平费率
	MarketTakeBound *big.Float `json:"marketTakeBound"` //  市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
}

type FilerType string

const (

	// "filterType": "PRICE_FILTER", // 价格限制
	// "maxPrice": "300", // 价格上限, 最大价格
	// "minPrice": "0.0001", // 价格下限, 最小价格
	// "tickSize": "0.0001" // 订单最小价格间隔
	PRICE_FILTER FilerType = "PRICE_FILTER"

	// "filterType": "LOT_SIZE", // 数量限制
	// "maxQty": "10000000", // 数量上限, 最大数量
	// "minQty": "1", // 数量下限, 最小数量
	// "stepSize": "1" // 订单最小数量间隔
	LOT_SIZE FilerType = "LOT_SIZE"

	// "filterType": "MARKET_LOT_SIZE", // 市价订单数量限制
	// "maxQty": "590119", // 数量上限, 最大数量
	// "minQty": "1", // 数量下限, 最小数量
	// "stepSize": "1" // 允许的步进值
	MARKET_LOT_SIZE FilerType = "MARKET_LOT_SIZE"

	// "filterType": "MAX_NUM_ORDERS", // 最多订单数限制
	// "limit": 200
	MAX_NUM_ORDERS FilerType = "MAX_NUM_ORDERS"

	// "filterType": "MAX_NUM_ALGO_ORDERS", // 最多条件订单数限制
	// "limit": 100
	MAX_NUM_ALGO_ORDERS FilerType = "MAX_NUM_ALGO_ORDERS"

	// "filterType": "MIN_NOTIONAL",  // 最小名义价值
	// "notional": "5.0",
	MIN_NOTIONAL FilerType = "MIN_NOTIONAL"

	// "filterType": "PERCENT_PRICE", // 价格比限制
	// "multiplierUp": "1.1500", // 价格上限百分比
	// "multiplierDown": "0.8500", // 价格下限百分比
	// "multiplierDecimal": 4
	PERCENT_PRICE FilerType = "PERCENT_PRICE"
)

type filter struct {
	FilterType FilerType `json:"filterType"` // 价格限制
	MaxPrice   string    `json:"maxPrice"`   // 价格上限, 最大价格
	MinPrice   string    `json:"minPrice"`   // 价格下限, 最小价格
	TickSize   string    `json:"tickSize"`   // 订单最小价格间隔

	MaxQty   string `json:"maxQty"`   // 数量上限, 最大数量
	MinQty   string `json:"minQty"`   // 数量下限, 最小数量
	StepSize string `json:"stepSize"` // 订单最小数量间隔

	Limit    string `json:"limit"`    //
	Notional string `json:"notional"` //

	MultiplierUp      string `json:"multiplierUp"`      // 价格比限制
	MultiplierDown    string `json:"multiplierDown"`    // 价格上限百分比
	MultiplierDecimal string `json:"multiplierDecimal"` // 价格下限百分比
}
