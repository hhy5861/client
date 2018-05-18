package client

import (
	"testing"
	"fmt"
)

var (
	urls = []string{"http://127.0.0.1:8090"}

	remoteList = map[string][]string{
		"kyc_serve": urls,
	}
)

func Test_Get(t *testing.T) {
	clients := NewClient(remoteList)

	var parms Params

	resp := clients.Get("kyc_serve", "v1/basic-info/search", parms, 0)

	fmt.Println(resp.GetMessage())
}
