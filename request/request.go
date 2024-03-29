package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Request struct {
	Path         string
	Params       ReqParams
	Method       RequestMethod
	Header       map[string]string
	ContentType  RequestContentType
	Payload      *Payload
	RawQuery     string
	BasicAuthKey string
}

type Payload struct {
	Body        *bytes.Buffer
	ContentType string
}

type RequestOptions func(*Request)

type ReqParams map[string]interface{}

func DefaultRequest(req *Request) {
	req.Method = GET
	req.ContentType = Json
}

func SetMethod(method RequestMethod) RequestOptions {
	return func(req *Request) {
		req.Method = method
	}
}

func SetContentType(content RequestContentType) RequestOptions {
	return func(req *Request) {
		req.ContentType = content
	}
}

func SetParams(Params ReqParams) RequestOptions {
	return func(req *Request) {
		req.Params = Params
	}
}

func SetHeader(Header map[string]string) RequestOptions {
	return func(req *Request) {
		req.Header = Header
	}
}

func SetRawQuery(RawQuery string) RequestOptions {
	return func(req *Request) {
		req.RawQuery = RawQuery
	}
}

// So far only string data has been parsed
// The following analysis is based on actual needs
// !!! Null data is filtered
func SortParams(params ReqParams) string {
	sorted := []string{}
	for k, v := range params {
		switch v := v.(type) {
		case string:
			if v != "" {
				sorted = append(sorted, k+"="+v)
			}
		}
	}
	return strings.Join(sorted, "&")
}

// Transform any structure into a standard transmission data format
// Currently, null data filtering is not supported
func Struct2Params(params interface{}) (Params ReqParams, err error) {
	b, err := json.Marshal(&params)
	if err != nil {
		return
	}
	Params = ReqParams{}
	err = json.Unmarshal(b, &Params)
	return
}

func NewRequest(Path string, Options ...RequestOptions) *Request {
	req := &Request{
		Path: Path,
	}

	DefaultRequest(req)

	for _, op := range Options {
		op(req)
	}

	return req
}

// Simplifying external writing of excessive code adds syntactic sugar
func (R *Request) Post(Params ReqParams) *Request {
	R.Method = POST
	R.Params = Params
	return R
}

func (R *Request) File(Payload *Payload) *Request {
	R.Post(nil)
	R.Payload = Payload
	return R
}

func (R *Request) BasicAuth(BasicAuth string) *Request {
	R.BasicAuthKey = BasicAuth
	return R
}

func (R *Request) Resources() (Body io.ReadCloser, err error) {
	resp, err := R.newRequest()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	Body = resp.Body
	return
}

func (R *Request) Send() (result []byte, err error) {
	resp, err := R.newRequest()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	return
}

func (R *Request) newRequest() (result *http.Response, err error) {
	client := &http.Client{}
	var req *http.Request
	var body *strings.Reader
	apiUrl := R.Path
	switch R.Method {
	case GET:
		body = strings.NewReader("")
		params := url.Values{}
		if R.Params != nil {
			for k, v := range R.Params {
				params.Add(k, fmt.Sprintf("%s", v))
			}
		}
		apiUrl += "?" + params.Encode()
	default:
		params, err := json.Marshal(R.Params)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(params))
	}
	if R.Payload != nil {
		log.Println("[Request]", "["+R.Method+"]", apiUrl)
		req, err = http.NewRequest(string(R.Method), apiUrl, R.Payload.Body)
		req.Header.Set("Content-Type", R.Payload.ContentType)
	} else if R.RawQuery != "" {
		log.Println("[Request]", "["+R.Method+"]", apiUrl, R.RawQuery)
		req, err = http.NewRequest(string(R.Method), apiUrl, nil)
		req.URL.RawQuery = R.RawQuery
	} else {
		log.Println("[Request]", "["+R.Method+"]", apiUrl, body)
		req, err = http.NewRequest(string(R.Method), apiUrl, body)
	}
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", string(R.ContentType))

	if R.Header != nil {
		for k, v := range R.Header {
			req.Header.Add(k, v)
		}
	}
	if R.BasicAuthKey != "" {
		req.Header.Add("Authorization", "Basic "+R.BasicAuthKey)
	}
	return client.Do(req)
}

func CreateFileData(Field, Filepath string, params ReqParams) (payload *Payload, err error) {
	Body := &bytes.Buffer{}
	writer := multipart.NewWriter(Body)
	file, err := os.Open(Filepath)
	if err != nil {
		return
	}
	defer file.Close()
	part1, err := writer.CreateFormFile(Field, filepath.Base(Filepath))
	_, err = io.Copy(part1, file)
	if err != nil {
		return
	}
	for k, v := range params {
		switch v := v.(type) {
		case string:
			if v != "" {
				_ = writer.WriteField(k, v)
			}
		}
	}
	err = writer.Close()
	if err != nil {
		return
	}
	payload = &Payload{
		Body:        Body,
		ContentType: writer.FormDataContentType(),
	}
	return
}
