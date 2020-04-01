package template

import "github.com/prodbox/weasn/kernel/message"

type IndustryClass struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type Template struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type GetIndustry struct {
	message.JsonErrorMessage
	PrimaryIndustry IndustryClass `json:"primary_industry"`
	DeputyIndustry  IndustryClass `json:"deputy_industry"`
}

type SendResponse struct {
	message.JsonErrorMessage
	MsgId string `json:"msgid"`
}
