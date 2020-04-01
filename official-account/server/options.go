package server

type Option func(*guard)

// EventHandler 事件处理
func EventHandler(handler IEvent) Option {
	return func(guard *guard) {
		guard.eventHandler(handler)
	}
}

// MessageHandler 消息处理
func MessageHandler(handler IMessage) Option {
	return func(guard *guard) {
		guard.messageHandler(handler)
	}
}
