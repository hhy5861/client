package client

import (
	"encoding/json"
)

type (
	Response struct {
		Body        string
		ResolveData *Resolve
		code        int
		message     string
		data        interface{}
	}

	Resolve struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func NewResponse(body string) *Response {
	res := &Response{
		Body: body,
	}

	res.ResolveBody()
	return res
}

func (r *Response) ResolveBody() *Response {
	var res Resolve

	err := json.Unmarshal([]byte(r.Body), &res)
	if err == nil {
		r.code = res.Code
		r.message = res.Message
		r.data = res.Data

		return nil
	}

	return r
}

func (r *Response) GetCode() int {
	return r.code
}

func (r *Response) GetData() interface{} {
	return r.data
}

func (r *Response) GetMessage() string {
	return r.message
}

func (r *Response) GetStruct(data interface{}) error {
	return json.Unmarshal([]byte(r.Body), &data)
}
