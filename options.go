package weasn

import (
	"github.com/prodbox/weasn/kernel/cache"
	"github.com/prodbox/weasn/kernel/context"
)

func AppId(v string) context.Option {
	return func(o *context.Options) {
		o.AppId = v
	}
}

func Secret(v string) context.Option {
	return func(o *context.Options) {
		o.Secret = v
	}
}

func Token(v string) context.Option {
	return func(o *context.Options) {
		o.Token = v
	}
}

func AESKey(v string) context.Option {
	return func(o *context.Options) {
		o.AESKey = v
	}
}

// 自定义缓存
func Cache(c cache.Cache) context.Option {
	return func(o *context.Options) {
		o.Cache = c
	}
}
