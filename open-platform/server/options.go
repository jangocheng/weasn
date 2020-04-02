package server

type Option func(*guard)

// EventHandler 事件处理
func Handler(handler IHandler) Option {
	return func(guard *guard) {
		guard.onHandler(handler)
	}
}
