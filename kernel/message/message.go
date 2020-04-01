package message

import "github.com/prodbox/weasn/kernel/message/internal"

type Message interface {
	Response() internal.Response
	Customer() internal.Customer
}

// NewText 文本消息
func NewText(v string) Message {
	return internal.NewText(v)
}

// NewImage 图片消息
func NewImage(mediaId string) Message {
	return internal.NewImage(mediaId)
}
