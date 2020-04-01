package open_platform

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account/server"
)

type Application struct {
	opts context.Options
	pool *context.Pool
}

func New(opts ...context.Option) *Application {
	app := &Application{
		opts: context.NewOptions(opts...).InitOptions(defaultOptions()),
	}

	app.pool = context.New(func() context.Context {
		return app.allocateContext()
	})
	return app
}

// defaultOptions 默认服务
func defaultOptions() []context.Option {
	return []context.Option{
		defaultHttpClient(),
		defaultCache(),
		defaultEncrypter(),
		defaultAccessToken(),
	}
}

// 初始化上下文
func (this *Application) allocateContext() context.Context {
	return context.NewContext(this.opts)
}

// Server 服务端
func (app *Application) Server(opts ...server.Option) server.Server {
	return server.New(app.pool).InitOptions(opts...)
}

func (this *Application) AccessToken() context.AccessToken {
	return this.opts.AccessToken
}
