package trait

import (
	"sync"
)

type Container struct {
	singleton sync.Map
}

// 单例
func (this *Container) Singleton(name string, fn func() interface{}) interface{} {
	var value interface{}
	if v, loaded := this.singleton.LoadOrStore(name, func() interface{} {
		value = fn()
		return value
	}()); loaded == true {
		return v
	}
	return value
}
