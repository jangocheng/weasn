package server

import "net/http"

// HandlerFunc 处理函数
type HandlerFunc func(mixed Mixed)

// Handler 消息监听接口
type IHandler interface {
	Authorized(mixed Mixed)
	UnAuthorized(mixed Mixed)
	UpdateAuthorized(mixed Mixed)
	ComponentVerifyTicket(mixed Mixed)
	NotifyThirdFasteregister(mixed Mixed)
}

// Server 服务端接口
type Server interface {
	Authorized(HandlerFunc) Server
	UnAuthorized(HandlerFunc) Server
	UpdateAuthorized(HandlerFunc) Server
	ComponentVerifyTicket(HandlerFunc) Server
	NotifyThirdFasteregister(HandlerFunc) Server
	ServeHTTP(http.ResponseWriter, *http.Request)
}
