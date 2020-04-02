package official

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/service"
	"github.com/prodbox/weasn/official-account"
	"github.com/prodbox/weasn/open-platform/authorizer/auth"
)

func New(opts context.Options, authorizerAppId, refreshToken string) *official_account.Application {
	return official_account.New(
		official_account.AppId(authorizerAppId),
		official_account.Token(opts.Token),
		official_account.AccessToken(service.NewAccessToken(auth.New(opts, authorizerAppId, refreshToken))),
		official_account.Encryptor(opts.Encrypter),
	)
}
