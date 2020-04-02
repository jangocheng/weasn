package base

import (
	"fmt"
	"net/url"

	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
)

type Client struct {
	*service.BaseClient
}

func New(opts context.Options) *Client {
	return &Client{BaseClient: service.NewBaseClient(opts)}
}

// 使用授权码换取公众号的授权信息
func (this *Client) HandleAuthorize(authCode string) (*QueryAuth, error) {
	params := map[string]string{
		"component_appid":    this.Options().AppId,
		"authorization_code": authCode,
	}
	var response QueryAuth
	return &response, this.Post("cgi-bin/component/api_query_auth", params, &response)
}

// 获取授权方的账户信息
func (this *Client) GetAuthorizer(authorizerAppId string) (*GetAuthorizerInfo, error) {
	params := map[string]string{
		"component_appid":  this.Options().AppId,
		"authorizer_appid": authorizerAppId,
	}
	var response GetAuthorizerInfo
	return &response, this.Post("cgi-bin/component/api_get_authorizer_info", params, &response)
}

// 获取授权方的选项设置信息
func (this *Client) GetAuthorizerOption(authorizerAppId, name string) (*GetAuthorizerOption, error) {
	params := map[string]string{
		"component_appid":  this.Options().AppId,
		"authorizer_appid": authorizerAppId,
		"option_name":      name,
	}
	var response GetAuthorizerOption
	return &response, this.Post("cgi-bin/component/api_get_authorizer_option", params, &response)
}

// 设置授权方的选项信息
func (this *Client) SetAuthorizerOption(authorizerAppId, name, value string) error {
	params := map[string]string{
		"component_appid":  this.Options().AppId,
		"authorizer_appid": authorizerAppId,
		"option_name":      name,
		"option_value":     value,
	}
	return this.Post("cgi-bin/component/api_set_authorizer_option", params, &message.JsonErrorMessage{})
}

// 拉取所有已授权的帐号基本信息
func (this *Client) GetAuthorizers(offset, count int) (*AuthorizerList, error) {
	params := map[string]interface{}{
		"component_appid": this.Options().AppId,
		"offset":          offset,
		"count":           count,
	}
	var response AuthorizerList
	return &response, this.Post("cgi-bin/component/api_get_authorizer_list", params, &response)
}

// GetPreAuthorizationUrl 获取用户授权页 URL
func (this *Client) GetPreAuthorizationUrl(redirectUri string) (string, error) {
	response, err := this.createPreAuthorizationCode()
	if err != nil {
		return "", err
	}

	queries := url.Values{}
	queries.Add("component_appid", this.Options().AppId)
	queries.Add("redirect_uri", redirectUri)
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?pre_auth_code=%s&%s", response.PreAuthCode, queries.Encode()), nil
}

func (this *Client) createPreAuthorizationCode() (*PreAuthCode, error) {
	params := map[string]string{
		"component_appid": this.Options().AppId,
	}
	var response PreAuthCode
	return &response, this.Post("cgi-bin/component/api_create_preauthcode", params, &response)
}

func (this *Client) ClearQuota() error {
	params := map[string]interface{}{
		"component_appid": this.Options().AppId,
	}
	return this.Post("cgi-bin/component/clear_quota", params, &message.JsonErrorMessage{})
}
