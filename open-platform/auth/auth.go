package auth

func New(ticket *VerifyTicket) *AccessToken {
	return &AccessToken{VerifyTicket: ticket}
}

type AccessToken struct {
	*VerifyTicket
}

func (this *AccessToken) TokenKey() string {
	return "component_access_token"
}

func (this *AccessToken) QueryName() string {
	return this.TokenKey()
}

func (this *AccessToken) Method() string {
	return "POST"
}

func (this *AccessToken) Endpoint() string {
	return "cgi-bin/component/api_component_token"
}

func (this *AccessToken) Credentials() map[string]string {
	return map[string]string{
		"component_appid":         this.opts.AppId,
		"component_appsecret":     this.opts.Secret,
		"component_verify_ticket": this.GetTicket(),
	}
}
