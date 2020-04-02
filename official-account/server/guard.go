package server

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/kernel/trait"
)

type guard struct {
	trait.Observable
	pool *context.Pool
}

func New(pool *context.Pool) *guard {
	return &guard{pool: pool}
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

func (this *guard) InitOptions(opts ...Option) *guard {
	for _, o := range opts {
		o(this)
	}
	return this
}

// Dispatch 分发事件
func (this *guard) Dispatch(payload []byte) interface{} {
	var mixed Mixed
	if err := xml.Unmarshal(payload, &mixed); err != nil {
		log.Println(err)
		return nil
	}

	condition := mixed.MsgType
	if condition == MsgTypeEvent {
		condition = mixed.EventType
	}

	// 调用监听事件
	if message := this.Observable.Dispatch(condition, mixed); message != nil {
		response := message.Response()
		response.To(mixed.FromUserName)
		response.From(mixed.ToUserName)
		response.MessageId(mixed.MsgId)
		response.CreatedAt(time.Now().Unix())
		return response
	}
	return nil
}

func (s *guard) ShouldReturnRawResponse(r *http.Request) []byte {
	if raw := r.URL.Query().Get("echostr"); len(raw) > 0 {
		fmt.Println(raw)
		return []byte(raw)
	}
	return nil
}

func (this *guard) onEventHandler(handler IEvent) {
	this.handle(EventSubscribe, handler.Subscribe)
	this.handle(EventUnsubscribe, handler.Unsubscribe)
}

func (this *guard) onMessageHandler(handler IMessage) {
	this.handle(MsgTypeText, handler.Text)
	this.handle(MsgTypeImage, handler.Image)
	this.handle(MsgTypeVoice, handler.Voice)
	this.handle(MsgTypeVideo, handler.Video)
	this.handle(MsgTypeLink, handler.Link)
	this.handle(MsgTypeEvent, handler.Event)
	this.handle(MsgTypeLocation, handler.Location)
	this.handle(MsgTypeShortVideo, handler.ShortVideo)
}

func (this *guard) handle(condition string, h HandlerFunc) Server {
	this.On(condition, func(payload interface{}) message.Message {
		return h(payload.(Mixed))
	})
	return this
}
