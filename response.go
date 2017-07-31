package htttp

import (
	"net/http"
)

type Response struct {
	Code int
	Data *interface{}
}

func Success(data interface{}) *Response {
	return &Response{
		Data: &data,
	}
}

func Status(code int) *Response {
	return &Response{
		Code: code,
	}
}

func (res *Response) Response(data interface{}) *Response {
	if res.Data != nil {
		panic("Overwriting response data")
	}

	res.Data = &data
	return res
}

func Error(code int) *Response {
	return Status(code).Response(http.StatusText(code))
}
