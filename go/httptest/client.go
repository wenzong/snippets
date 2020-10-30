package demo

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.mch.weixin.qq.com"
)

type Client struct {
	c   *http.Client
	url string
}

func NewClient(c *http.Client, URL string) *Client {
	return &Client{c: c, url: URL}
}

func (mc *Client) UnifiedOrder(req *Request) (*Response, error) {
	b, _ := xml.Marshal(req)
	r, _ := mc.c.Post(mc.url+"/pay/unifiedorder", "application/xml", bytes.NewBuffer(b))
	defer r.Body.Close()

	// TODO: Check res.StatusCode

	response := &Response{}
	b, _ = ioutil.ReadAll(r.Body)
	_ = xml.Unmarshal(b, &response)

	return response, nil
}
