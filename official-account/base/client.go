package base

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/kernel/service"
)

type Client struct {
	baseClient *service.BaseClient
}

func New(opts context.Options) *Client {
	return &Client{baseClient: service.NewBaseClient(opts)}
}

// clearQuota 清理接口调用次数
func (this *Client) ClearQuota() error {
	params := map[string]string{
		"appid": this.baseClient.Options().AppId,
	}
	return this.baseClient.Post("cgi-bin/clear_quota", params, &message.JsonErrorMessage{})
}

// getValidIps 获取微信服务器 IP (或IP段)
func (this *Client) GetValidIps() []string {
	var ips struct {
		message.JsonErrorMessage
		IpList []string `json:"ip_list"`
	}
	this.baseClient.Get("cgi-bin/getcallbackip", nil, &ips)
	return ips.IpList
}
