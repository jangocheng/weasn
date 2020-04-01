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
	// 用户
	User() *user.UserClient
	// 基础接口
	Base() *base.Client
	// 临时素材
	Media() *media.Client
	// 网页授权
	OAuth() *oauth.Client
	// 微信AccessToken
	AccessToken() context.AccessToken
	// 客服
	CustomerService() *customer.Client
	// 模板消息
	TemplateMessage() *template.Client
	// 服务端
	Server(...server.Option) server.Server
}

func NewOfficialAccount(opts ...context.Option) OfficialAccount {
	return official_account.New(opts...)
}
