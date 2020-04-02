package auth

import (
	"fmt"
	"github.com/prodbox/weasn/kernel/context"
)

type auth struct {
	opts            context.Options
	authorizerAppId string
	refreshToken    string
}

func New(opts context.Options, authorizerAppId, refreshToken string) *auth {
	return &auth{
		opts:            opts,
		authorizerAppId: authorizerAppId,
		refreshToken:    refreshToken,
	}
}

func (this *auth) Options() context.Options {
	return this.opts
}

func (self *auth) TokenKey() string {
	return "authorizer_access_token"
}

func (self *auth) QueryName() string {
	return "access_token"
}

func (this *auth) Method() string {
	return "POST"
}

func (this *auth) Endpoint() string {
	componentAccessToken, _ := this.opts.AccessToken.GetToken()
	return fmt.Sprintf("cgi-bin/component/api_authorizer_token?component_access_token=%s", componentAccessToken)
}

func (this *auth) Credentials() map[string]string {
	return map[string]string{
		"component_appid":          this.opts.AppId,
		"authorizer_appid":         this.authorizerAppId,
		"authorizer_refresh_token": this.refreshToken,
	}
}
