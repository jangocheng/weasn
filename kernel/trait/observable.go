package trait

import (
	"github.com/prodbox/weasn/kernel/message"
	"sync"
)

type HandlerFunc func(payload interface{}) message.Message

// 观察者
type Observable struct {
	handlers sync.Map
}

// 注册一个监听事件
func (this *Observable) On(condition string, handler HandlerFunc) {
	this.handlers.Store(condition, handler)
}

// 获取所有监听事件
func (this *Observable) GetHandlers() map[string]HandlerFunc {
	handlers := make(map[string]HandlerFunc)
	this.handlers.Range(func(k, v interface{}) bool {
		handlers[k.(string)] = v.(HandlerFunc)
		return true
	})
	return handlers
}

// 分发
func (this *Observable) Dispatch(condition string, payload interface{}) message.Message {
	if handler, ok := this.handlers.Load(condition); ok == true {
		return handler.(HandlerFunc)(payload)
	}
	return nil
}
