package official_account

import (
	"github.com/go-resty/resty/v2"
	"github.com/prodbox/weasn/kernel/cache"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/official-account/auth"
	"time"
)

func defaultHttpClient() context.Option {
	return func(o *context.Options) {
		o.HttpClient = resty.New().SetHostURL("https://api.weixin.qq.com/").SetTimeout(5 * time.Second)
	}
}

func defaultCache() context.Option {
	return func(o *context.Options) {
		if o.Cache == nil {
			if err := cache.DefaultCache.Start(); err != nil {
				panic(err)
			}
			o.Cache = cache.DefaultCache
		}
	}
}

// 加密/解密模块
func defaultEncrypter() context.Option {
	return func(o *context.Options) {
		if o.Encrypter == nil {
			o.Encrypter = service.NewEncryptor(o.AppId, o.Token, o.AESKey)
		}
	}
}

func defaultAccessToken() context.Option {
	return func(o *context.Options) {
		if o.AccessToken == nil {
			o.AccessToken = service.NewAccessToken(auth.New(o.Clone()))
		}
	}
}

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

// 加密/解密模块
func AccessToken(token context.AccessToken) context.Option {
	return func(o *context.Options) {
		o.AccessToken = token
	}
}
