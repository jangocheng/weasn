package service

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/kernel/message"
)

type BaseClient struct {
	opts context.Options
}

func NewBaseClient(opts context.Options) *BaseClient {
	return &BaseClient{opts: opts}
}

// Options 获取初始化参数
func (this *BaseClient) Options() context.Options {
	return this.opts
}

func (this *BaseClient) Get(endpoint string, params map[string]string, resp message.JsonError) error {
	return this.request("GET", endpoint, resp, func(r *resty.Request) *resty.Request {
		return r.SetQueryParams(params)
	})
}

func (this *BaseClient) Post(endpoint string, body interface{}, resp message.JsonError) error {
	return this.request("POST", endpoint, resp, func(r *resty.Request) *resty.Request {
		return r.SetBody(body)
	})
}

func (this *BaseClient) Upload(endpoint string, files map[string]string, params map[string]string, resp message.JsonError) error {
	return this.request("POST", endpoint, resp, func(r *resty.Request) *resty.Request {
		return r.SetQueryParams(params).SetFiles(files)
	})
}

func (this *BaseClient) request(method, endpoint string, resp message.JsonError, beforeFn func(r *resty.Request) *resty.Request) error {
	params, err := this.opts.AccessToken.GetTokenParams()
	if err != nil {
		return err
	}

	// 请求微信
	response, err := beforeFn(this.opts.HttpClient.R()).SetHeader("Accept", "application/json").SetQueryParams(params).Execute(method, endpoint)
	if err != nil {
		return err
	}

	if response.Header().Get("Content-Type") != "application/json" {
		return json.Unmarshal(response.Body(), resp)
	}
	return nil
}
