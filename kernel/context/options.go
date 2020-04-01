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
