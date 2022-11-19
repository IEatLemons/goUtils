package ethereum

import (
	"encoding/json"
	"errors"

	"github.com/IEatLemons/goUtils/communicate"
)

const (
	EtherscanUrl string = "https://api.etherscan.io/api"
)

var EtherScan *Etherscan

type Etherscan struct {
	Apikey string
}

type EtherscanModule string
type EtherscanAction string
type EtherscanTag string
type EtherscanSort string

const (
	Account  EtherscanModule = "account"
	Contract EtherscanModule = "contract"
	Logs     EtherscanModule = "logs"
	Token    EtherscanModule = "token"
)

const (
	Balance         EtherscanAction = "balance"
	Txlist          EtherscanAction = "txlist"
	GetLogs         EtherscanAction = "getLogs"
	TokenHolderList EtherscanAction = "tokenholderlist"
)

const (
	Earliest EtherscanTag = "earliest"
	Pending  EtherscanTag = "pending"
	Latest   EtherscanTag = "latest"
)

const (
	ASC  EtherscanSort = "ASC"
	DESC EtherscanSort = "DESC"
)

type BaseParams struct {
	Module EtherscanModule
	Action EtherscanAction
}

func InitEtherscan(Apikey string) *Etherscan {
	if EtherScan == nil {
		EtherScan = &Etherscan{
			Apikey: Apikey,
		}
	}
	return EtherScan
}

func NewEtherscan() (*Etherscan, error) {
	if EtherScan == nil {
		return nil, errors.New("not instantiated")
	}
	return EtherScan, nil
}

func (E *Etherscan) Request(Params Request, Result interface{}) (err error) {
	JsonBytes, err := communicate.Request(communicate.GET, EtherscanUrl, Params.GetParams(E.Apikey), communicate.FormUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(JsonBytes), Result)
	return
}
