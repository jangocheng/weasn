package user

import "github.com/prodbox/weasn/kernel/message"

type UserInfo struct {
	message.JsonErrorMessage
	OpenId         string  `json:"openid"`
	UnionId        string  `json:"unionid"`
	Nickname       string  `json:"nickname"`
	Sex            uint    `json:"sex"`
	Language       string  `json:"language"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Country        string  `json:"country"`
	HeadImgUrl     string  `json:"headimgurl"`
	Remark         string  `json:"remark"`
	GroupId        uint    `json:"groupid"`
	TagidList      []int64 `json:"tagid_list"`
	Subscribe      uint    `json:"subscribe"`
	SubscribeTime  int64   `json:"subscribe_time"`
	SubscribeScene string  `json:"subscribe_scene"`
	QRScene        int64   `json:"qr_scene"`
	QRSceneStr     string  `json:"qr_scene_str"`
}

type UserGet struct {
	message.JsonErrorMessage
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}
