package server

import (
	"encoding/xml"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/kernel/trait"
	"github.com/prodbox/weasn/open-platform/auth"
	"net/http"
	"reflect"
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
		// 此处可以写日志
	}
}

// Dispatch 分发事件
func (this *guard) Dispatch(payload []byte) interface{} {
	var mixed Mixed
	if err := xml.Unmarshal(payload, &mixed); err != nil {
		return nil
	}

	// 更新ComponentVerifyTicket
	if mixed.InfoType == EVENT_COMPONENT_VERIFY_TICKET {
		this.SetTicket(mixed.ComponentVerifyTicket)
	}

	// 调用监听事件
	this.Observable.Dispatch(mixed.InfoType, mixed)
	return nil
}

// ShouldReturnRawResponse
func (this *guard) ShouldReturnRawResponse(r *http.Request) []byte {
	return nil
}

// Authorized 处理授权成功事件
func (this *guard) Authorized(fn HandlerFunc) Server {
	return this.handle(EVENT_AUTHORIZED, fn)
}

func (this *guard) UnAuthorized(fn HandlerFunc) Server {
	return this.handle(EVENT_UNAUTHORIZED, fn)
}

func (this *guard) UpdateAuthorized(fn HandlerFunc) Server {
	return this.handle(EVENT_UPDATE_AUTHORIZED, fn)
}

func (this *guard) ComponentVerifyTicket(fn HandlerFunc) Server {
	return this.handle(EVENT_COMPONENT_VERIFY_TICKET, fn)
}

func (this *guard) NotifyThirdFasteregister(fn HandlerFunc) Server {
	return this.handle(EVENT_NOTIFY_THIRD_FASTEREGISTER, fn)
}

func (this *guard) bindObject(object interface{}) {
	v := reflect.ValueOf(object)
	t := v.Type()

	for i := 0; i < v.NumMethod(); i++ {
		itemFunc := v.Method(i).Interface().(HandlerFunc)
		if v, ok := objectMapper[t.Method(i).Name]; ok {
			this.handle(v, itemFunc)
		}
	}
}

func (this *guard) handle(condition string, h HandlerFunc) Server {
	this.On(condition, func(payload interface{}) message.Message {
		h(payload.(Mixed))
		return nil
	})
	return this
}
