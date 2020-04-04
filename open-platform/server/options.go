package server

type Option func(*guard)

// EventHandler 事件处理
func BindEventObject(object IEventObject) Option {
	return func(guard *guard) {
		guard.bindObject(object)
	}
}
