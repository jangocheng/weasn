package server

type Option func(*guard)

// BindEventObject
func BindEventObject(object IEventObject) Option {
	return func(guard *guard) {
		guard.bindObject(object)
	}
}

// BindMessageObject
func BindMessageObject(object IMessageObject) Option {
	return func(guard *guard) {
		guard.bindObject(object)
	}
}
