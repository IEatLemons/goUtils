package ethereum

type UnifyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AccountTxlistResponse struct {
	UnifyResponse
	Result []struct {
		BlockNumber       int    `json:"blockNumber,string"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		TransactionIndex  string `json:"transactionIndex"`
		From              string `json:"from"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		IsError           string `json:"isError"`
		TxreceiptStatus   string `json:"txreceipt_status"`
		Input             string `json:"input"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		GasUsed           string `json:"gasUsed"`
		Confirmations     string `json:"confirmations"`
		MethodId          string `json:"methodId"`
		FunctionName      string `json:"functionName"`
	} `json:"result"`
}

type LogsGetLogsResponse struct {
	UnifyResponse
	Result []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		TimeStamp        string   `json:"timeStamp"`
		GasPrice         string   `json:"gasPrice"`
		GasUsed          string   `json:"gasUsed"`
		LogIndex         string   `json:"logIndex"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
	}
}

type TokenHolderListResponse struct {
	UnifyResponse
	Result []struct {
		TokenHolderAddress  string `json:"TokenHolderAddress"`
		TokenHolderQuantity string `json:"TokenHolderQuantity"`
	}
}
