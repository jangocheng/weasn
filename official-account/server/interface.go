package server

import (
	"github.com/prodbox/weasn/kernel/message"
	"net/http"
)

// HandlerFunc 处理函数
type HandlerFunc func(Mixed) message.Message

// 事件通知接口
type IEventObject interface {
	Scan(Mixed) message.Message
	Click(Mixed) message.Message
	View(Mixed) message.Message
	Location(Mixed) message.Message
	Subscribe(Mixed) message.Message
	Unsubscribe(Mixed) message.Message
	ScancodePush(Mixed) message.Message
	ScancodeWaitmsg(Mixed) message.Message
	PicSysphoto(Mixed) message.Message
	PicPhotoOrAlbum(Mixed) message.Message
	PicWeixin(Mixed) message.Message
	LocationSelect(Mixed) message.Message
	TemplateSendJobFinish(Mixed) message.Message
}

// 消息通知接口
type IMessageObject interface {
	Text(Mixed) message.Message
	Image(Mixed) message.Message
	Voice(Mixed) message.Message
	Video(Mixed) message.Message
	Music(Mixed) message.Message
	News(Mixed) message.Message
	Link(Mixed) message.Message
	Transfer(Mixed) message.Message
	ShortVideo(Mixed) message.Message
	Location(Mixed) message.Message
}

// Server 服务端接口
type Server interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
