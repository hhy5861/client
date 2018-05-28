package client

import (
	"time"
)

type (
	Client struct {
		remotes map[string][]string
		request *request
		debug   bool
	}

	Params map[string]string
)

var (
	client *Client
)

// init client
func NewClient(remotes map[string][]string, debug bool) *Client {
	client = &Client{
		remotes: remotes,
		debug:   debug,
	}

	return client
}

// get client
func GetClient() *Client {
	return client
}

//request get
func (c *Client) Get(
	remote,
	path string,
	queryParams Params) *Response {

	c.request = NewRequest(c.remotes, c.debug)

	return c.request.SetRemote(remote).SetPath(path).SetParam(queryParams).Get()
}

//request post
func (c *Client) Post(
	remote,
	path string,
	dataForm Params) *Response {

	c.request = NewRequest(c.remotes, c.debug)

	return c.request.SetRemote(remote).SetPath(path).SetParam(dataForm).Post()
}

//request post json
func (c *Client) PostJson(
	remote,
	path string,
	dataJson Params) *Response {

	c.request = NewRequest(c.remotes, c.debug)

	return c.request.SetRemote(remote).SetPath(path).SetParam(dataJson).PostJson()
}

func (c *Client) Delete(
	remote,
	path string,
	dataForm Params) *Response {

	c.request = NewRequest(c.remotes, c.debug)

	return c.request.SetRemote(remote).SetPath(path).SetParam(dataForm).Delete()
}

// set request header data params
func (c *Client) SetHeader(data map[string]string) *Client {
	c.request.SetHeader(data)

	return c
}

// set request time out params
func (c *Client) SetTimeOut(times time.Duration) *Client {
	c.SetTimeOut(times)

	return c
}
