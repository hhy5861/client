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

func NewClient(remotes map[string][]string, debug bool) *Client {
	return &Client{
		remotes: remotes,
		debug:   debug,
	}
}

//request get
func (c *Client) Get(
	remote,
	path string,
	query Params) *Response {

	c.request = NewRequest(c.remotes, c.debug)

	return c.request.SetRemote(remote).SetPath(path).SetParam(query).Get()
}

func (c *Client) Post(remote, path string, formParams map[string]interface{}, timeout time.Duration, async bool) {

}

func (c *Client) PostJson(remote, path string, data interface{}, timeout time.Duration, async bool) {

}

func (c *Client) SetHeader(data map[string]string) *Client {
	c.request.SetHeader(data)

	return nil
}

func (c *Client) SetTimeOut(times time.Duration) {

}
