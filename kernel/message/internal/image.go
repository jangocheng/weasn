package internal

type Image struct {
	MsgType string
	MediaId string
}

type ResponseImage struct {
	ResponseMessage
	Image struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Image"`
}

type CustomerImage struct {
	CustomerMessage
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(mediaId string) *Image {
	return &Image{
		MsgType: "image",
		MediaId: mediaId,
	}
}

func (this *Image) Response() Response {
	text := new(ResponseImage)
	text.Image.MediaId = this.MediaId
	text.MsgType = CDATA("image")
	return text
}

func (this *Image) Customer() Customer {
	text := new(CustomerImage)
	text.Image.MediaId = this.MediaId
	text.MsgType = this.MsgType
	return text
}
