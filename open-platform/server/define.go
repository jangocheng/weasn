package server

import "encoding/xml"

const (
	// 处理授权成功事件
	EVENT_AUTHORIZED = "authorized"

	// 处理授权取消事件
	EVENT_UNAUTHORIZED = "unauthorized"

	// 处理授权更新事件
	EVENT_UPDATE_AUTHORIZED = "updateauthorized"

	// 处理微信10分钟推送一次component_verify_ticket
	EVENT_COMPONENT_VERIFY_TICKET = "component_verify_ticket"
)

// Mixed 微信每隔10分钟发过来一个数据ticket|授权成功通知|取消授权通知|授权更新通知
type Mixed struct {
	XMLName                      xml.Name `xml:"xml"`
	AppId                        string   `xml:"AppId"`                        //第三方平台appid
	CreateTime                   int64    `xml:"CreateTime"`                   //时间戳
	InfoType                     string   `xml:"InfoType"`                     //component_verify_ticket
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"`        //Ticket内容
	AuthorizerAppId              string   `xml:"AuthorizerAppid"`              //公众号appid
	AuthorizationCode            string   `xml:"AuthorizationCode"`            //授权码
	AuthorizationCodeExpiredTime string   `xml:"AuthorizationCodeExpiredTime"` //过期时间
	PreAuthCode                  string   `xml:"PreAuthCode"`                  //预授权码
}
