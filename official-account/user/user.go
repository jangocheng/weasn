package user

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
)

type UserClient struct {
	baseClient *service.BaseClient
}

func NewUser(opts context.Options) *UserClient {
	return &UserClient{baseClient: service.NewBaseClient(opts)}
}

// Get 获取用户信息(单个)
func (this *UserClient) Get(openId string) (*UserInfo, error) {
	params := map[string]string{
		"openid": openId,
		"lang":   "zh_CN",
	}

	var response UserInfo
	return &response, this.baseClient.Get("cgi-bin/user/info", params, &response)
}

// Get 获取用户信息(多个)
func (this *UserClient) Select(openIds []string) ([]UserInfo, error) {
	params := make([]map[string]string, 0)
	for _, v := range openIds {
		params = append(params, map[string]string{"openid": v, "lang": "zh_CN"})
	}

	var response struct {
		message.JsonErrorMessage
		List []UserInfo `json:"user_info_list"`
	}
	return response.List, this.baseClient.Post("cgi-bin/user/info", map[string]interface{}{"user_list": params}, &response)
}

// List 获取用户列表
func (this *UserClient) List(nextOpenId ...string) (*UserGet, error) {
	params := map[string]string{}
	if len(nextOpenId) > 0 {
		params["next_openid"] = nextOpenId[0]
	}
	var response UserGet
	return &response, this.baseClient.Get("cgi-bin/user/get", params, &response)
}

// Remark 修改用户备注
func (this *UserClient) Remark(openId, remark string) error {
	params := map[string]string{
		"openid": openId,
		"remark": remark,
	}
	return this.baseClient.Post("cgi-bin/user/info/updateremark", params, &message.JsonErrorMessage{})
}

// Block 拉黑用户
func (this *UserClient) Block(openidList ...string) error {
	params := make([]string, 0)
	for _, v := range openidList {
		params = append(params, v)
	}
	return this.baseClient.Post("cgi-bin/tags/members/batchblacklist", map[string][]string{"openid_list": params}, &message.JsonErrorMessage{})
}

// UnBlock 取消拉黑
func (this *UserClient) UnBlock(openidList ...string) error {
	params := make([]string, 0)
	for _, v := range openidList {
		params = append(params, v)
	}
	return this.baseClient.Post("cgi-bin/tags/members/batchunblacklist", map[string][]string{"openid_list": params}, &message.JsonErrorMessage{})
}

// BlackList 获取黑名单
func (this *UserClient) BlackList(beginOpenid ...string) (*UserGet, error) {
	params := map[string]string{}
	if len(beginOpenid) > 0 {
		params["begin_openid"] = beginOpenid[0]
	}
	var response UserGet
	return &response, this.baseClient.Get("cgi-bin/tags/members/getblacklist", params, &response)
}
