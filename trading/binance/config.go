package binance

const (
	recvWindow = 5000
)

type SpotAPIURL string
type SpotWSURL string
type SpotPATH string

const (
	// https://api.binance.com
	// https://api-gcp.binance.com
	// https://api1.binance.com
	// https://api2.binance.com
	// https://api3.binance.com
	// https://api4.binance.com
	MainnetSpotAPI SpotAPIURL = "https://api.binance.com"
	TestnetSpotAPI SpotAPIURL = "https://testnet.binance.vision"
	MainnetSpotWSS SpotWSURL  = "wss://stream.binance.com:9443/ws"
	TestnetSpotWSS SpotWSURL  = "wss://testnet.binance.vision/ws"
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

type Environment string

const (
	Mainnet Environment = "mainnet"
	Testnet Environment = "testnet"
)

func NewBinance(env Environment, path SpotPATH) *binance {
	var api SpotAPIURL
	var wss SpotWSURL
	if env == Mainnet {
		api = MainnetSpotAPI
		wss = MainnetSpotWSS
	} else {
		api = MainnetSpotAPI
		wss = MainnetSpotWSS
	}
	return &binance{
		SpotApi:  api,
		SpotWss:  wss,
		SpotPath: path,
	}
}

type binance struct {
	SpotApi  SpotAPIURL
	SpotWss  SpotWSURL
	SpotPath SpotPATH
}

func (b *binance) getSpotApiUrl() string {
	return string(b.SpotApi) + string(b.SpotPath)
}