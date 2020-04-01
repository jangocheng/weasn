package official_account

import (
	"github.com/prodbox/weasn/basic/media"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account/base"
	"github.com/prodbox/weasn/official-account/customer"
	"github.com/prodbox/weasn/official-account/oauth"
	"github.com/prodbox/weasn/official-account/server"
	"github.com/prodbox/weasn/official-account/template"
	"github.com/prodbox/weasn/official-account/user"
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

// Base 基础接口
func (this *Application) Base() *base.Client {
	return base.New(this.opts)
}

// Media 临时素材
func (this *Application) Media() *media.Client {
	return media.New(this.opts)
}

// OAuth 网页授权
func (this *Application) OAuth() *oauth.Client {
	return oauth.New(this.opts)
}

// User 用户
func (app *Application) User() *user.UserClient {
	return user.NewUser(app.opts)
}

func (this *Application) AccessToken() context.AccessToken {
	return this.opts.AccessToken
}

// CustomerService 客服
func (this *Application) CustomerService() *customer.Client {
	return customer.New(this.opts)
}

// TemplateMessage 模板消息
func (this *Application) TemplateMessage() *template.Client {
	return template.New(this.opts)
}
