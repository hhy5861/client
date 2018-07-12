package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
	"encoding/json"
	"net/http"
)

type (
	request struct {
		remote     string
		url        string
		SuperAgent *gorequest.SuperAgent
		remotes    map[string][]string
		timeOut    time.Duration
		param      Params
		header     map[string]string
	}
)

var (
	t       *Utils
	TimeOut time.Duration = 3
)

func init() {
	t = NewUtil()
}

func NewRequest(remote map[string][]string, debug bool) *request {

	return &request{
		remotes:    remote,
		SuperAgent: gorequest.New().SetDebug(debug),
	}
}

func (r *request) SetTimeOut(times time.Duration) {
	r.timeOut = times
}

func (r *request) GetTimeOut() time.Duration {
	if r.timeOut <= 0 {
		r.timeOut = TimeOut
	}

	return r.timeOut * time.Second
}

func (r *request) SetRemote(remote string) *request {
	remoteArray, ok := r.remotes[remote]
	if ok {
		max := len(remoteArray)
		num := t.GenerateRangeNum(0, max)
		r.remote = remoteArray[num]
	}

	return r
}

func (r *request) SetPath(path string) *request {
	path = strings.TrimRight(strings.TrimLeft(path, "/"), "/")

	r.url = fmt.Sprintf("%s/%s",
		strings.Trim(r.remote, "/"),
		path)

	return r
}

func (r *request) SetHeader(data map[string]string) *request {
	r.header = data

	return r
}

func (r *request) SetParam(param Params) *request {
	r.param = param

	return r
}

func (r *request) GetParam() *request {
	for k, v := range r.param {
		r.SuperAgent.Param(k, v)
	}

	return r
}

func (r *request) GetHeader() *request {
	i := 0
	for k, v := range r.header {
		if i > 0 {
			r.SuperAgent.AppendHeader(k, v)
		} else {
			r.SuperAgent.Set(k, v)
		}

		i++
	}

	return r
}

func (r *request) Get() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Get(r.url)

	r.GetParam()

	res, body, err := r.SuperAgent.End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) Post() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Post(r.url)

	r.GetHeader().GetParam()

	res, body, err := r.SuperAgent.End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) PostUrlEncode() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Post(r.url)

	r.GetHeader()

	res, body, err := r.SuperAgent.Send(r.param).End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) Put() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Put(r.url)

	r.GetParam()

	res, body, err := r.SuperAgent.End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) PostJson() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Post(r.url)

	paramsJson, errMsg := json.Marshal(r.param)
	if errMsg != nil {
		return NewResponse("{}", 406)
	}

	res, body, err := r.SuperAgent.Send(paramsJson).End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) Delete() *Response {
	r.SuperAgent.Timeout(r.GetTimeOut()).Delete(r.url)

	r.GetParam()

	res, body, err := r.SuperAgent.End()
	if err == nil {
		return r.responseCode(body, res.StatusCode)
	}

	return NewResponse(body, http.StatusBadGateway)
}

func (r *request) responseCode(body string, statusCode int) *Response {

	if statusCode <= 206 && statusCode >= 200 {
		return NewResponse(body, http.StatusOK)
	}

	return NewResponse(body, statusCode)
}
