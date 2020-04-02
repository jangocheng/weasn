package server

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/kernel/trait"
	"github.com/prodbox/weasn/open-platform/auth"
)

type guard struct {
	trait.Observable
	*auth.VerifyTicket
	pool *context.Pool
}

func New(pool *context.Pool, ticket *auth.VerifyTicket) *guard {
	return &guard{VerifyTicket: ticket, pool: pool}
}

func (this *guard) InitOptions(opts ...Option) *guard {
	for _, o := range opts {
		o(this)
	}
	return this
}

// ServeHTTP 服务端入口
func (this *guard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := this.pool.Acquire(w, r)
	defer this.pool.Release(ctx)

	if err := service.NewServerGuard(this, ctx).Serve(); err != nil {
		log.Println(err)
		// 此处可以写日志
	}
}

// Dispatch 分发事件
func (this *guard) Dispatch(payload []byte) interface{} {
	var mixed Mixed
	if err := xml.Unmarshal(payload, &mixed); err != nil {
		log.Println(err)
		return nil
	}

	// 更新ComponentVerifyTicket
	if mixed.InfoType == EVENT_COMPONENT_VERIFY_TICKET {
		this.SetTicket(mixed.ComponentVerifyTicket)
	}
	log.Println("Dispatch ======= >")
	// 调用监听事件
	this.Observable.Dispatch(mixed.InfoType, mixed)
	return nil
}

// ShouldReturnRawResponse
func (this *guard) ShouldReturnRawResponse(r *http.Request) []byte {
	return nil
}

// Authorized 处理授权成功事件
func (this *guard) Authorized(h HandlerFunc) Server {
	return this.handle(EVENT_AUTHORIZED, h)
}

// UnAuthorized 处理授权取消事件
func (this *guard) UnAuthorized(h HandlerFunc) Server {
	return this.handle(EVENT_UNAUTHORIZED, h)
}

// UpdateAuthorized 处理授权更新事件
func (this *guard) UpdateAuthorized(h HandlerFunc) Server {
	return this.handle(EVENT_UPDATE_AUTHORIZED, h)
}

// VerifyTicket 处理微信10分钟推送一次的 component_verify_ticket
func (this *guard) ComponentVerifyTicket(h HandlerFunc) Server {
	return this.handle(EVENT_COMPONENT_VERIFY_TICKET, h)
}

// NotifyThirdFasteregister 快速创建小程序的信息
func (this *guard) NotifyThirdFasteregister(h HandlerFunc) Server {
	return this.handle("notify_third_fasteregister", h)
}

func (this *guard) onHandler(h IHandler) {
	this.Authorized(h.Authorized)
	this.UnAuthorized(h.UnAuthorized)
	this.UpdateAuthorized(h.UpdateAuthorized)
	this.ComponentVerifyTicket(h.ComponentVerifyTicket)
	this.NotifyThirdFasteregister(h.NotifyThirdFasteregister)
}

// 添加监听事件
func (this *guard) handle(condition string, h HandlerFunc) Server {
	this.On(condition, func(payload interface{}) message.Message {
		h(payload.(Mixed))
		return nil
	})
	return this
}
