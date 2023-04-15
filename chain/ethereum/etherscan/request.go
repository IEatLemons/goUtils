package etherscan

import (
	"encoding/json"

	"github.com/IEatLemons/goUtils/communicate"
)

type Request interface {
	GetParams(Apikey string) communicate.ReqParams
}

func getParams(Request interface{}, Apikey string) (Params communicate.ReqParams) {
	Params = communicate.ReqParams{}
	inrec, _ := json.Marshal(Request)
	json.Unmarshal(inrec, &Params)
	Params["apikey"] = Apikey
	return
}

type UnifyRequest struct {
	Module EtherscanModule `json:"module"`
	Action EtherscanAction `json:"action"`
}

type Page struct {
	Page   string `json:"page"`
	Offset string `json:"offset"`
	Sort   string `json:"sort"`
}

type TxListRequest struct {
	UnifyRequest
	Page
	Address    string `json:"address"`
	Startblock string `json:"startblock"`
	Endblock   string `json:"endblock"`
}

func (R *TxListRequest) GetParams(Apikey string) communicate.ReqParams {
	return getParams(R, Apikey)
}

type GetLogsRequest struct {
	UnifyRequest
	Page
	Address   string `json:"address"`
	Topic0    string `json:"topic0"`
	FromBlock string `json:"fromBlock"`
	ToBlock   string `json:"toBlock"`
}

func (R *GetLogsRequest) GetParams(Apikey string) communicate.ReqParams {
	return getParams(R, Apikey)
}

type TokenHolderListRequest struct {
	UnifyRequest
	Page
	ContractAddress string `json:"contractaddress"`
}

func (R *TokenHolderListRequest) GetParams(Apikey string) communicate.ReqParams {
	return getParams(R, Apikey)
}
