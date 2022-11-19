package communicate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Method string

const (
	POST Method = "POST"
	GET  Method = "GET"
	PUT  Method = "PUT"
)

type ReqParams map[string]string

type ContentType string

const (
	Json    ContentType = "application/json"
	FormUrl ContentType = "application/x-www-form-urlencoded"
)

func SortParams(params ReqParams) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		sorted[i] = k + "=" + url.QueryEscape(params[k])
		i++
	}
	return strings.Join(sorted, "&")
}

func Request(method Method, path string, params ReqParams, ContentType ContentType) (result string, err error) {
	client := &http.Client{}
	sorted := ""
	switch ContentType {
	case Json:
		by, _ := json.Marshal(params)
		sorted = string(by)
	default:
		sorted = SortParams(params)
	}
	var req *http.Request
	if method == GET {
		req, _ = http.NewRequest(string(method), path+"?"+sorted, strings.NewReader(""))
	} else {
		req, _ = http.NewRequest(string(method), path, strings.NewReader(sorted))
	}
	req.Header.Add("Content-Type", string(ContentType))

	fmt.Println("Client Do......", path)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}
