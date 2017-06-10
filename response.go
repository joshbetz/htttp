package htttp

type Response struct {
	Code int
	Data *interface{}
}

func Res(data interface{}) *Response {
	return &Response{
		Data: &data,
	}
}

func ResWithCode(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: &data,
	}
}

func Error(code int, data interface{}) *Response {
	return ResWithCode(code, data)
}
