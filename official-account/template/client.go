package template

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/service"
)

type Client struct {
	baseClient *service.BaseClient
}

func New(opts context.Options) *Client {
	return &Client{baseClient: service.NewBaseClient(opts)}
}
