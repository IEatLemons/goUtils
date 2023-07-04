package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"

	"github.com/IEatLemons/goUtils/request"
)

func (b *binance) QuestSimple(values *url.Values) ([]byte, error) {
	return request.NewRequest(b.getSpotApiUrl(), []request.RequestOptions{
		request.SetRawQuery(values.Encode()),
	}...).Send()
}

func (b *binance) QuestGet(values *url.Values) ([]byte, error) {
	b.basic(values)
	return request.NewRequest(b.getSpotApiUrl(), []request.RequestOptions{
		request.SetRawQuery(values.Encode()),
	}...).Send()
}

func (b *binance) QuestPost(values *url.Values) ([]byte, error) {
	b.basic(values)
	return request.NewRequest(b.getSpotApiUrl(), []request.RequestOptions{
		request.SetRawQuery(values.Encode()),
	}...).Post(nil).Send()
}

func (b *binance) HmacGet(apiKey, secretKey string, values *url.Values) ([]byte, error) {
	valStr := b.signature(apiKey, secretKey, values)
	options := b.signatureOptions(apiKey, valStr)
	return b.quest(options...)
}

func (b *binance) HmacPost(apiKey, secretKey string, values *url.Values) ([]byte, error) {
	valStr := b.signature(apiKey, secretKey, values)
	options := b.signatureOptions(apiKey, valStr)
	options = append(options, request.SetMethod(request.POST))
	options = append(options, request.SetContentType(request.FormUrl))
	return b.quest(options...)
}

func (b *binance) basic(values *url.Values) {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	values.Set("timestamp", timestamp)
	values.Set("recvWindow", strconv.Itoa(recvWindow))
}

func (b *binance) quest(Options ...request.RequestOptions) ([]byte, error) {
	return request.NewRequest(b.getSpotApiUrl(), Options...).Send()
}

func (b *binance) signature(apiKey, secretKey string, values *url.Values) (valStr string) {
	b.basic(values)
	values.Set("signature", b.generateSignature(values.Encode(), secretKey))
	valStr = values.Encode()
	return
}

func (b *binance) generateSignature(query string, secretKey string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(query))
	return hex.EncodeToString(mac.Sum(nil))
}

func (b *binance) signatureOptions(apiKey, valStr string) []request.RequestOptions {
	return []request.RequestOptions{
		request.SetHeader(map[string]string{
			"X-MBX-APIKEY": apiKey,
		}),
		request.SetRawQuery(valStr),
	}
}
