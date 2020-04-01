package server

import (
	"github.com/prodbox/weasn/kernel/message"
	"net/http"
)

// HandlerFunc 处理函数
type HandlerFunc func(Mixed) message.Message

// 事件通知接口
type IEvent interface {
	Subscribe(Mixed) message.Message
	Unsubscribe(Mixed) message.Message
}

// 消息通知接口
type IMessage interface {
	Text(Mixed) message.Message
	Image(Mixed) message.Message
	Voice(Mixed) message.Message
	Video(Mixed) message.Message
	Link(Mixed) message.Message
	Event(Mixed) message.Message
	ShortVideo(Mixed) message.Message
	Location(Mixed) message.Message
}

// Server 服务端接口
type Server interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
