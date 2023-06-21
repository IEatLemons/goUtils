package binance

type PATH string

const (
	APIURL = "https://api.binance.com"
	WSURL  = "wss://stream.binance.com:9443/ws"
)

type Method string

const (
	SUBSCRIBE Method = "SUBSCRIBE"
)

type Interval string

const (
	OneDay         Interval = "1d"
	OneHours       Interval = "1h"
	FourHours      Interval = "4h"
	OneMinutes     Interval = "1m"
	FifteenMinutes Interval = "15m"
	OneSecond      Interval = "1s"
)
