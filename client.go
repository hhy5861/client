package client

import (
	"time"
)

type (
	client struct {
		remotes map[string][]string
		debug   bool
	}

	Params map[string]string
)

var (
	c *client
	r *request
)

// init client
func NewClient(remotes map[string][]string, debug bool) *client {
	c = &client{
		remotes: remotes,
		debug:   debug,
	}

	return c
}

// get client
func GetClient() *client {
	if r == nil {
		r = NewRequest(c.remotes, c.debug)
	}

	return c
}

//request get
func (c *client) Get(
	remote,
	path string,
	queryParams interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(queryParams).Get()
}

//request post
func (c *client) Post(
	remote,
	path string,
	dataForm interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(dataForm).Post()
}

//request post
func (c *client) PostUrlEncode(
	remote,
	path string,
	dataForm interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(dataForm).PostUrlEncode()
}

//request put
func (c *client) Put(
	remote,
	path string,
	dataForm interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(dataForm).Put()
}

//request post json
func (c *client) PostJson(
	remote,
	path string,
	dataJson interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(dataJson).PostJson()
}

func (c *client) Delete(
	remote,
	path string,
	dataForm interface{}) *Response {

	return r.SetRemote(remote).SetPath(path).SetParam(dataForm).Delete()
}

// set request header data params
func (c *client) SetHeader(data map[string]string) *client {
	r.SetHeader(data)

	return c
}

// set request time out params
func (c *client) SetTimeOut(times time.Duration) *client {
	r.SetTimeOut(times)

	return c
}

func (c *client) AddParams(key, value string) {
	r.SuperAgent.QueryData.Add(key, value)
}
