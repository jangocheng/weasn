package internal

import (
	"encoding/xml"
)

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

/******  客服消息  ********/
type CustomService struct {
	KfAccount string `json:"kf_account"`
}

type CustomerMessage struct {
	ToUser        string         `json:"touser"`
	MsgType       string         `json:"msgtype"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

func (this *CustomerMessage) To(to string) {
	this.ToUser = to
}

func (this *CustomerMessage) KfAccount(kfAccount string) {
	this.CustomService = &CustomService{KfAccount: kfAccount}
}

/******* 服务端响应消息 ********/
type ResponseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	MsgType      CDATA    `xml:"MsgType"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgId        int64    `xml:"MsgId"`
}

func (this *ResponseMessage) To(to string) {
	this.ToUserName = CDATA(to)
}

func (this *ResponseMessage) From(from string) {
	this.FromUserName = CDATA(from)
}

func (this *ResponseMessage) MessageId(msgId int64) {
	this.MsgId = msgId
}

func (this *ResponseMessage) CreatedAt(ctime int64) {
	this.CreateTime = ctime
}
