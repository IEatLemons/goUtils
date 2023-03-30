package openai

import (
	"encoding/json"
	"errors"

	"github.com/IEatLemons/goUtils/request"
)

type Turbo struct {
	*OPENAI
	Path string
}

type TurboResp struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created uint64          `json:"created"`
	Choices []*TurboChoices `json:"choices"`
	Usage   *TurboUsage     `json:"usage"`
	Error   *TurboError     `json:"error"`
}

type TurboChoices struct {
	Index        uint64        `json:"index"`
	Message      *TurboMessage `json:"message"`
	FinishReason string        `json:"finish_reason"`
}

type TurboMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type TurboUsage struct {
	PromptTokens     uint64 `json:"prompt_tokens"`
	CompletionTokens uint64 `json:"completion_tokens"`
	TotalTokens      uint64 `json:"total_tokens"`
}

type TurboError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   interface{}
	Code    interface{}
}

type Role string

const (
	SYSTEM Role = "system"
	USER   Role = "user"
	ASSIS  Role = "assistant"
)

func (OPENAI *OPENAI) Turbo() *Turbo {
	return &Turbo{
		OPENAI: OPENAI,
		Path:   "/v1/chat/completions",
	}
}

func (T *Turbo) Chat(Msg []*TurboMessage) (respMsg *TurboMessage, err error) {
	resp, err := request.NewRequest(HOST+T.Path, []request.RequestOptions{
		request.SetHeader(map[string]string{
			"Authorization": "Bearer " + T.Key,
		}),
	}...).Post(request.ReqParams{
		"model":       TURBO,
		"messages":    Msg,
		"temperature": 1,
	}).Send()
	if err != nil {
		return
	}

	var r TurboResp
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	if r.Error != nil {
		err = errors.New(r.Error.Message)
		return
	}
	respMsg = r.Choices[0].Message
	return
}
