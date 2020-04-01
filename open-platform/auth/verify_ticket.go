package auth

import (
	"fmt"
	"github.com/prodbox/weasn/kernel/context"
)

type VerifyTicket struct {
	opts context.Options
}

func NewVerifyTicket(opts context.Options) *VerifyTicket {
	return &VerifyTicket{opts: opts}
}

func (this *VerifyTicket) Options() context.Options {
	return this.opts
}

// SetTicket 设置 verifyticket
func (this *VerifyTicket) SetTicket(v string) error {
	return this.opts.Cache.Set(this.getCacheKey(), v, 3600)
}

// GetTicket 获取缓存Ticket
func (this *VerifyTicket) GetTicket() string {
	// 验证缓存
	if exist := this.opts.Cache.IsExist(this.getCacheKey()); exist == true {
		if v := this.opts.Cache.Get(this.getCacheKey()); v != nil {
			return v.(string)
		}
	}
	return ""
}

// getCacheKey 缓存key
func (this *VerifyTicket) getCacheKey() string {
	return fmt.Sprintf("weasn.open_platform.verify_ticket.%s", this.opts.AppId)
}
