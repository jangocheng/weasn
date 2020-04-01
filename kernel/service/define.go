package service

import "encoding/xml"

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// 微信加密消息固定格式
type EncryptRquest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"-"` // 开发者微信号
	Encrypt      string   // 加密的消息报文
	MsgSignature string   // 报文签名
	TimeStamp    string   // 时间戳
	Nonce        string   // 随机数
}

// 加密消息(响应)
type EncryptResponse struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATA    `xml:"Encrypt"`
	MsgSignature CDATA    `xml:"MsgSignature"`
	TimeStamp    int64    `xml:"TimeStamp"`
	Nonce        CDATA    `xml:"Nonce"`
}
