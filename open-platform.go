package weasn

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account/server"
	"github.com/prodbox/weasn/open-platform"
)

type OpenPlatform interface {
	Server(opts ...server.Option) server.Server
	AccessToken() context.AccessToken
}

func NewOpenPlatform(opts ...context.Option) OpenPlatform {
	return open_platform.New(opts...)
}
