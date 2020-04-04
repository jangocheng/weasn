package context

import (
	"github.com/go-resty/resty/v2"
	"github.com/prodbox/weasn/kernel/cache"
)

type Option func(*Options)

type Options struct {
	AppId  string
	Secret string
	Token  string
	AESKey string
	// 全局缓存组件
	Cache cache.Cache
	// 网络请求
	HttpClient *resty.Client
	// 全局加密器
	Encrypter Encrypter
	// AccessToken
	AccessToken AccessToken
}

func NewOptions(opts ...Option) Options {
	opt := Options{}
	return opt.InitOptions(opts)
}

// InitOptions 初始化参数
func (opt Options) InitOptions(opts []Option) Options {
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func (o Options) Clone() Options {
	return Options{
		AppId:       o.AppId,
		Secret:      o.Secret,
		Token:       o.Token,
		AESKey:      o.AESKey,
		Cache:       o.Cache,
		Encrypter:   o.Encrypter,
		HttpClient:  o.HttpClient,
		AccessToken: o.AccessToken,
	}
}

func AppId(v string) Option {
	return func(o *Options) {
		o.AppId = v
	}
}

func Secret(v string) Option {
	return func(o *Options) {
		o.Secret = v
	}
}

func Token(v string) Option {
	return func(o *Options) {
		o.Token = v
	}
}

func AESKey(v string) Option {
	return func(o *Options) {
		o.AESKey = v
	}
}

// 加密/解密模块
func WithEncrypter(e Encrypter) Option {
	return func(o *Options) {
		o.Encrypter = e
	}
}

// 加密/解密模块
func WithAccessToken(token AccessToken) Option {
	return func(o *Options) {
		o.AccessToken = token
	}
}
