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

// 加密/解密模块
func Encryptor(e context.Encrypter) context.Option {
	return func(o *context.Options) {
		o.Encrypter = e
	}
}

// 自定义缓存
func Cache(c cache.Cache) context.Option {
	return func(o *context.Options) {
		o.Cache = c
	}
}
