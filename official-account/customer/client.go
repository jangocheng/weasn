package customer

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

// List 获取所有客服
func (this *Client) List() ([]KFBasicInfo, error) {
	var response struct {
		message.JsonErrorMessage
		List []KFBasicInfo `json:"kf_list"`
	}
	return response.List, this.baseClient.Get("cgi-bin/customservice/getkflist", nil, &response)
}

// Online 获取所有在线的客服
func (this *Client) Online() ([]OnlineKFBasicInfo, error) {
	var response struct {
		message.JsonErrorMessage
		List []OnlineKFBasicInfo `json:"kf_online_list"`
	}
	return response.List, this.baseClient.Get("cgi-bin/customservice/getonlinekflist", nil, &response)
}

// Create 添加客服
func (this *Client) Create(account, nickname string) error {
	params := map[string]string{
		"f_account": account,
		"nickname":  nickname,
	}
	return this.baseClient.Post("customservice/kfaccount/add", params, &message.JsonErrorMessage{})
}

// Update 修改客服
func (this *Client) Update(account, nickname string) error {
	params := map[string]string{
		"f_account": account,
		"nickname":  nickname,
	}
	return this.baseClient.Post("customservice/kfaccount/update", params, &message.JsonErrorMessage{})
}

// Delete 删除账号
func (this *Client) Delete(account string) error {
	params := map[string]string{
		"f_account": account,
	}
	return this.baseClient.Post("customservice/kfaccount/del", params, &message.JsonErrorMessage{})
}

// Invite 邀请微信用户加入客服
func (this *Client) Invite(account, wechatId string) error {
	params := map[string]string{
		"f_account": account,
		"invite_wx": wechatId,
	}
	return this.baseClient.Post("customservice/kfaccount/inviteworker", params, &message.JsonErrorMessage{})
}

// SetAvatar 设置客服头像
func (this *Client) SetAvatar(account, path string) error {
	params := map[string]string{
		"kf_account": account,
	}
	files := map[string]string{
		"media": path,
	}
	return this.baseClient.Upload("customservice/kfaccount/uploadheadimg", files, params, &message.JsonErrorMessage{})
}

// Message 主动发送消息给用户
func (this *Client) Message(message message.Message) *messenger {
	return &messenger{client: this, message: message}
}

// Messages 获取客服与客户聊天记录
func (this *Client) Messages(startTime, endTime, msgId, number int64) ([]MsgRecord, error) {
	params := map[string]int64{
		"starttime": startTime,
		"endtime":   endTime,
		"msgid":     msgId,
		"number":    number,
	}

	var response struct {
		message.JsonErrorMessage
		List []MsgRecord `json:"recordlist"`
	}
	return response.List, this.baseClient.Post("customservice/msgrecord/getmsglist", params, &response)
}

// send 发送客服消息
func (this *Client) send(msg interface{}) error {
	return this.baseClient.Post("customservice/msgrecord/getmsglist", msg, &message.JsonErrorMessage{})
}
