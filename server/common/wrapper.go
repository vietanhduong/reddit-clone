package common

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Content interface{} `json:"content,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
