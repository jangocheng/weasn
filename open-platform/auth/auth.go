package auth

import "github.com/prodbox/weasn/kernel/context"

type auth struct {
	*VerifyTicket
}

func New(ticket *VerifyTicket) *auth {
	return &auth{VerifyTicket: ticket}
}

func (this *auth) Options() context.Options {
	return this.opts
}

func (this *auth) TokenKey() string {
	return "component_access_token"
}

func (this *auth) QueryName() string {
	return this.TokenKey()
}

func (this *auth) Method() string {
	return "POST"
}

func (this *auth) Endpoint() string {
	return "cgi-bin/component/api_component_token"
}

func (this *auth) Credentials() map[string]string {
	return map[string]string{
		"component_appid":         this.opts.AppId,
		"component_appsecret":     this.opts.Secret,
		"component_verify_ticket": this.GetTicket(),
	}
}
