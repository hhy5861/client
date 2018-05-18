package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

type (
	request struct {
		remote     string
		url        string
		timeout    time.Duration
		SuperAgent *gorequest.SuperAgent
		remotes    map[string][]string
		timeOut    time.Duration
		param      Params
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
	for k, v := range data {
		r.SuperAgent.Set(k, v)
	}

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

func (r *request) Get() *Response {

	r.SuperAgent.Timeout(r.GetTimeOut()).Get(r.url)
	r.GetParam()
	res, body, err := r.SuperAgent.End()

	if err == nil && res.StatusCode >= 200 {

		return NewResponse(body)
	}

	return NewResponse("{}")
}
