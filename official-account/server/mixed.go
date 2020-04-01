package server

import "encoding/xml"

//EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

// 微信服务器推送过来的消息(事件)的合集.
type Mixed struct {
	XMLName      xml.Name `xml:"xml"`
	FromUserName string   `xml:"FromUserName"`
	ToUserName   string   `xml:"ToUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MsgId        int64    `xml:"MsgId"`
	Content      string   `xml:"Content"`
	MediaId      string   `xml:"MediaId"` // 消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string   `xml:"ThumbMediaId"`
	PicURL       string   `xml:"PicUrl"` // 图片链接
	LocationX    float64  `xml:"Location_X"`
	LocationY    float64  `xml:"Location_Y"`
	Scale        float64  `xml:"Scale"`
	Label        string   `xml:"Label"`
	Title        string   `xml:"Title"`
	Description  string   `xml:"Description"`
	URL          string   `xml:"Url"`
	EventType    string   `xml:"Event"`
	EventKey     string   `xml:"EventKey"`
	Ticket       string   `xml:"Ticket"`
	Latitude     string   `xml:"Latitude"`
	Longitude    string   `xml:"Longitude"`
	Precision    string   `xml:"Precision"`
	MenuID       string   `xml:"MenuId"`
	Status       string   `xml:"Status"`
	SessionFrom  string   `xml:"SessionFrom"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}
}
