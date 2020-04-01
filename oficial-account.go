package weasn

import (
	"github.com/prodbox/weasn/basic/media"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account"
	"github.com/prodbox/weasn/official-account/base"
	"github.com/prodbox/weasn/official-account/customer"
	"github.com/prodbox/weasn/official-account/server"
)

type OfficialAccount interface {
	Base() *base.Client
	Media() *media.Client
	AccessToken() context.AccessToken
	CustomerService() *customer.Client
	Server(...server.Option) server.Server
}

func NewOfficialAccount(opts ...context.Option) OfficialAccount {
	return official_account.New(opts...)
}
