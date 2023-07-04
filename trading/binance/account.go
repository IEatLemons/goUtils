package binance

import (
	"encoding/json"
	"log"
	"net/url"

	binance_struct "github.com/IEatLemons/goUtils/trading/binance/struct"
)

type account struct {
	account         *binance
	apiRestrictions *binance
	apiKey          string
	secretKey       string
}

func InitAccount(apiKey, secretKey string, env Environment) *account {
	return &account{
		account:         NewBinance(env, "/api/v3/account"),
		apiRestrictions: NewBinance(env, "/sapi/v1/account/apiRestrictions"),
		apiKey:          apiKey,
		secretKey:       secretKey,
	}
}

func (c *account) Account() (data *binance_struct.AccountInfo, err error) {
	result, err := c.account.HmacGet(c.apiKey, c.secretKey, &url.Values{})
	if err != nil {
		return
	}
	log.Println("["+c.account.SpotPath+"][result]", string(result))
	data = &binance_struct.AccountInfo{}
	err = json.Unmarshal(result, data)
	if err != nil {
		return
	}
	return
}

func (c *account) ApiRestrictions() (data *binance_struct.ApiRestrictions, err error) {
	result, err := c.apiRestrictions.HmacGet(c.apiKey, c.secretKey, &url.Values{})
	if err != nil {
		return
	}
	log.Println("["+c.apiRestrictions.SpotPath+"][result]", string(result))
	data = &binance_struct.ApiRestrictions{}
	err = json.Unmarshal(result, data)
	if err != nil {
		return
	}
	return
}
