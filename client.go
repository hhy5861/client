package client

import (
	"time"
)

type (
	Client struct {
		remotes map[string][]string
		debug   bool
	}

	Params map[string]string
)

var (
	c *Client
	r *request
)

// init client
func NewClient(remotes map[string][]string, debug bool) *Client {
	c = &Client{
		remotes: remotes,
		debug:   debug,
	}

	return c
}

// get client
func GetClient() *Client {
	getRequest()

	return c
}

func getRequest() *request {
	if r == nil {
		r = NewRequest(c.remotes, c.debug)
	}

	return r
}

//request get
func (c *Client) Get(
	remote,
	path string,
	queryParams interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(queryParams).Get()
}

//request post
func (c *Client) Post(
	remote,
	path string,
	dataForm interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(dataForm).Post()
}

//request post
func (c *Client) PostUrlEncode(
	remote,
	path string,
	dataForm interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(dataForm).PostUrlEncode()
}

//request put
func (c *Client) Put(
	remote,
	path string,
	dataForm interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(dataForm).Put()
}

//request post json
func (c *Client) PostJson(
	remote,
	path string,
	dataJson interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(dataJson).PostJson()
}

func (c *Client) Delete(
	remote,
	path string,
	dataForm interface{}) *Response {

	return getRequest().SetRemote(remote).SetPath(path).SetParam(dataForm).Delete()
}

// set request header data params
func (c *Client) SetHeader(data map[string]string) *Client {
	getRequest().SetHeader(data)

	return c
}

// set request time out params
func (c *Client) SetTimeOut(times time.Duration) *Client {
	getRequest().SetTimeOut(times)

	return c
}

func (c *Client) AddParams(key, value string) {
	getRequest().SuperAgent.QueryData.Add(key, value)
}
