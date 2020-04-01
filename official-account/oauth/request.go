package oauth

import (
	"net/url"
)

type request struct {
	appId string
	scope string
	state string
}

func (this *request) clone() *request {
	clone := *this
	return &clone
}

func (this *request) Scope(scope string) *request {
	this.scope = scope
	return this
}

func (this *request) State(state string) *request {
	this.state = state
	return this
}

func (this *request) Redirect(redirectUri string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + url.QueryEscape(this.appId) +
		"&redirect_uri=" + url.QueryEscape(redirectUri) +
		"&response_type=code&scope=" + url.QueryEscape(this.scope) +
		"&state=" + url.QueryEscape(this.state) +
		"#wechat_redirect"
}
