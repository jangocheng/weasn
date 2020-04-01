package internal

type Text struct {
	MsgType string
	Content string
}

type ResponseText struct {
	ResponseMessage
	Content CDATA `xml:"Content" json:"Content"`
}

type CustomerText struct {
	CustomerMessage
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(v string) *Text {
	return &Text{
		MsgType: "text",
		Content: v,
	}
}

func (this *Text) Response() Response {
	text := new(ResponseText)
	text.Content = CDATA(this.Content)
	text.MsgType = CDATA("text")
	return text
}

func (this *Text) Customer() Customer {
	text := new(CustomerText)
	text.MsgType = this.MsgType
	text.Text.Content = this.Content
	return text
}
