package openai

type OPENAI struct {
	Key       string
	BasicAuth string
}

const HOST = "https://api.openai.com"

func Init(key, BasicAuth string) *OPENAI {
	return &OPENAI{
		Key:       key,
		BasicAuth: BasicAuth,
	}
}
