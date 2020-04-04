package weasn

import (
	"github.com/prodbox/weasn/basic/media"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account"
	"github.com/prodbox/weasn/official-account/base"
	"github.com/prodbox/weasn/official-account/customer"
	"github.com/prodbox/weasn/official-account/oauth"
	"github.com/prodbox/weasn/official-account/server"
	"github.com/prodbox/weasn/official-account/template"
	"github.com/prodbox/weasn/official-account/user"
)

// 微信公众号
type OfficialAccount interface {
	// 用户(单例)
	User() *user.UserClient
	// 基础接口(单例)
	Base() *base.Client
	// 临时素材(单例)
	Media() *media.Client
	// 网页授权(单例)
	OAuth() *oauth.Client
	// 微信AccessToken(单例)
	AccessToken() context.AccessToken
	// 客服(单例)
	CustomerService() *customer.Client
	// 模板消息(单例)
	TemplateMessage() *template.Client
	// 服务端(单例)
	Server(...server.Option) server.Server
}

func NewOfficialAccount(opts ...context.Option) OfficialAccount {
	return official_account.New(opts...)
}
