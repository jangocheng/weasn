package auth

import "github.com/prodbox/weasn/kernel/context"

type auth struct {
	opts context.Options
}

func New(opts context.Options) *auth {
	return &auth{opts: opts}
}

func (this *auth) Options() context.Options {
	return this.opts
}

func (this *auth) TokenKey() string {
	return "access_token"
}

func (this *auth) QueryName() string {
	return this.TokenKey()
}

func (this *auth) Method() string {
	return "GET"
}

func (this *auth) Endpoint() string {
	return "cgi-bin/token"
}

func (this *auth) Credentials() map[string]string {
	return map[string]string{
		"grant_type": "client_credential",
		"appid":      this.opts.AppId,
		"secret":     this.opts.Secret,
	}
}
