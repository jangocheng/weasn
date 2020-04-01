package customer

import (
	"encoding/json"
	"fmt"
	"github.com/prodbox/weasn/kernel/message"
)

type messenger struct {
	to      string
	account string
	client  *Client
	message message.Message
}

// From 设置客服
func (this *messenger) From(account string) *messenger {
	this.account = account
	return this
}

// To 用户openid
func (this *messenger) To(openid string) *messenger {
	this.to = openid
	return this
}

// Send 发送
func (this *messenger) Send() error {
	message := this.message.Customer()
	message.To(this.to)
	if len(this.account) > 0 {
		message.KfAccount(this.account)
	}

	bytes, err := json.Marshal(message)
	fmt.Println(string(bytes), err)
	return this.client.send(message)
}
