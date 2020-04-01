package server

import "net/http"

// HandlerFunc 处理函数
type HandlerFunc func(mixed Mixed)

// Handler 消息监听接口
type MessageHandler interface {
	Authorized(HandlerFunc) Server
	UnAuthorized(HandlerFunc) Server
	UpdateAuthorized(HandlerFunc) Server
	VerifyTicket(HandlerFunc) Server
	NotifyThirdFasteregister(HandlerFunc) Server
}

// Server 服务端接口
type Server interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
