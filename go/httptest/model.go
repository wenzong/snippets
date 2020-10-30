package demo

import "encoding/xml"

type Request struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          string   `xml:"appid"`
	Attach         string   `xml:"attach"`
	Body           string   `xml:"body"`
	MchID          string   `xml:"mch_id"`
	Detail         string   `xml:"detail"`
	NonceStr       string   `xml:"nonce_str"`
	NotifyUrl      string   `xml:"notify_url"`
	OpenID         string   `xml:"openid"`
	OutTradeNo     string   `xml:"out_trade_no"`
	SpbillCreateIP string   `xml:"spbill_create_ip"`
	TotalFee       int      `xml:"total_fee"`
	TradeType      string   `xml:"trade_type"`
	Sign           string   `xml:"sign"`
}

type Response struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	AppID      string   `xml:"appid"`
	MchID      string   `xml:"mch_id"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	ResultCode string   `xml:"result_code"`
	PrepayID   string   `xml:"prepay_id"`
	TradeType  string   `xml:"trade_type"`
}
