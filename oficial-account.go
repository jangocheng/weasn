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
)

type OfficialAccount interface {
	Base() *base.Client
	Media() *media.Client
	OAuth() *oauth.Client
	AccessToken() context.AccessToken
	CustomerService() *customer.Client
	TemplateMessage() *template.Client
	Server(...server.Option) server.Server
}

func NewOfficialAccount(opts ...context.Option) OfficialAccount {
	return official_account.New(opts...)
}
