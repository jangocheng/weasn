package official_account

import (
	"github.com/prodbox/weasn/basic/media"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/trait"
	"github.com/prodbox/weasn/official-account/base"
	"github.com/prodbox/weasn/official-account/customer"
	"github.com/prodbox/weasn/official-account/oauth"
	"github.com/prodbox/weasn/official-account/server"
	"github.com/prodbox/weasn/official-account/template"
	"github.com/prodbox/weasn/official-account/user"
)

type Application struct {
	trait.Container
	opts context.Options
	pool *context.Pool
}

func New(opts ...context.Option) *Application {
	options := context.NewOptions(opts...).InitOptions(defaultOptions())

	app := &Application{opts: options}
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
func (this *Application) Server(opts ...server.Option) server.Server {
	return this.Singleton("server", func() interface{} {
		return server.New(this.pool).InitOptions(opts...)
	}).(server.Server)
}

// Base 基础接口
func (this *Application) Base() *base.Client {
	return this.Singleton("base", func() interface{} {
		return base.New(this.opts)
	}).(*base.Client)
}

// Media 临时素材
func (this *Application) Media() *media.Client {
	return this.Singleton("media", func() interface{} {
		return media.New(this.opts)
	}).(*media.Client)
}

// OAuth 网页授权
func (this *Application) OAuth() *oauth.Client {
	return this.Singleton("oauth", func() interface{} {
		return oauth.New(this.opts)
	}).(*oauth.Client)
}

// User 用户
func (this *Application) User() *user.UserClient {
	return this.Singleton("user", func() interface{} {
		return user.NewUser(this.opts)
	}).(*user.UserClient)
}

func (this *Application) AccessToken() context.AccessToken {
	return this.opts.AccessToken
}

// CustomerService 客服
func (this *Application) CustomerService() *customer.Client {
	return this.Singleton("customer_service", func() interface{} {
		return customer.New(this.opts)
	}).(*customer.Client)
}

// TemplateMessage 模板消息
func (this *Application) TemplateMessage() *template.Client {
	return this.Singleton("template_message", func() interface{} {
		return template.New(this.opts)
	}).(*template.Client)
}
