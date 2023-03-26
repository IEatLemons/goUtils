package request

type RequestContentType string

const (
	Json        RequestContentType = "application/json"
	FormUrl     RequestContentType = "application/x-www-form-urlencoded"
	MulFormData RequestContentType = "multipart/form-data"
)
