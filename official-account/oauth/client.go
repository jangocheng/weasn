package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"net/http"
)

type Client struct {
	request *request
	opts    context.Options
}

func New(opts context.Options) *Client {
	return &Client{opts: opts}
}

func (this *Client) User(code string) (*UserInfo, error) {
	if token, err := this.getAccessToken(code); err == nil {
		return this.getUserByToken(token)
	} else {
		return nil, err
	}
}

func (this *Client) Scope(scope string) *request {
	return this.clone().request.Scope(scope)
}

func (this *Client) State(state string) *request {
	return this.clone().request.State(state)
}

func (this *Client) GetRedirectURL(redirectUri string) string {
	return this.clone().request.Redirect(redirectUri)
}

func (this *Client) Redirect(w http.ResponseWriter, r *http.Request, redirectUri string) {
	http.Redirect(w, r, this.GetRedirectURL(redirectUri), 302)
}

func (this *Client) getAccessToken(code string) (*AccessToken, error) {
	params := map[string]string{
		"appid":      this.opts.AppId,
		"secret":     this.opts.Secret,
		"code":       code,
		"grant_type": "authorization_code",
	}

	var response AccessToken
	return &response, this.httpRequest("GET", "sns/oauth2/access_token", params, &response)
}

func (this *Client) getUserByToken(token *AccessToken) (*UserInfo, error) {
	params := map[string]string{
		"access_token": token.AccessToken,
		"openid":       token.OpenId,
		"lang":         "zh_CN",
	}

	var response UserInfo
	return &response, this.httpRequest("GET", "sns/userinfo", params, &response)
}

func (this *Client) httpRequest(method, endpoint string, params map[string]string, resp message.JsonError) error {
	response, err := this.opts.HttpClient.R().SetQueryParams(params).SetResult(resp).Execute(method, endpoint)
	if err != nil {
		return err
	}

	// 手动解析json
	if response.Header().Get("Content-Type") != "application/json" {
		if err := json.Unmarshal(response.Body(), resp); err != nil {
			return err
		}
	}

	if resp.GetCode() != 0 {
		return fmt.Errorf("errcode=%v , errmsg=%v", resp.GetCode(), resp.GetMsg())
	}

	return nil
}

func (this *Client) clone() *Client {
	oauth := &Client{opts: this.opts}
	if this.request == nil {
		oauth.request = &request{appId: this.opts.AppId, scope: "snsapi_userinfo", state: "STATE"}
	} else {
		oauth.request = this.request.clone()
	}
	return oauth
}
