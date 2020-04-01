package template

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

// setIndustry 修改账号所属行业
func (this *Client) SetIndustry(industryId1, industryId2 string) error {
	params := map[string]string{
		"industry_id1": industryId1,
		"industryId2":  industryId2,
	}
	return this.baseClient.Post("cgi-bin/template/api_set_industry", params, &message.JsonErrorMessage{})
}

// GetIndustry 获取支持的行业列表
func (this *Client) GetIndustry() (*GetIndustry, error) {
	var response GetIndustry
	return &response, this.baseClient.Get("cgi-bin/template/get_industry", nil, &response)
}

// AddTemplate 添加模板
func (this *Client) AddTemplate(shortId string) (string, error) {
	params := map[string]string{
		"template_id_short": shortId,
	}

	var response struct {
		message.JsonErrorMessage
		TemplateId string `json:"template_id"`
	}
	return response.TemplateId, this.baseClient.Post("gi-bin/template/api_add_template", params, &response)
}

// GetPrivateTemplates 获取所有模板列表
func (this *Client) GetPrivateTemplates() []Template {
	var response struct {
		message.JsonErrorMessage
		List []Template `json:"template_list"`
	}
	this.baseClient.Post("cgi-bin/template/get_all_private_template", nil, &response)
	return response.List
}

// DeletePrivateTemplate 删除模板
func (this *Client) DeletePrivateTemplate(templateId string) error {
	params := map[string]string{
		"template_id": templateId,
	}
	return this.baseClient.Post("cgi-bin/template/get_all_private_template", params, &message.JsonErrorMessage{})
}

func (this *Client) Send(data map[string]interface{}) (*SendResponse, error) {
	var response SendResponse
	return &response, this.baseClient.Post("cgi-bin/message/template/send", data, &response)
}
