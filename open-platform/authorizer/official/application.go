package official

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/official-account"
	"github.com/prodbox/weasn/open-platform/authorizer/auth"
)

func New(opts context.Options, authorizerAppId, refreshToken string) *official_account.Application {
	return official_account.New(
		context.AppId(authorizerAppId),
		context.Token(opts.Token),
		context.WithAccessToken(service.NewAccessToken(auth.New(opts, authorizerAppId, refreshToken))),
		context.WithEncrypter(opts.Encrypter),
	)
}
