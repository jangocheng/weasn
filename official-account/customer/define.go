package customer

// 客服基本信息
type KFBasicInfo struct {
	Account    string `json:"kf_account"`
	Nick       string `json:"kf_nick"`
	Id         string `json:"kf_id"`
	Headimgurl string `json:"kf_headimgurl"`
}

// 在线客服信息
type OnlineKFBasicInfo struct {
	Account      string `json:"kf_account"`
	Status       int    `json:"status"`
	Id           string `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

// 消息记录
type MsgRecord struct {
	Openid   string `json:"openid"`
	Opercode int64  `json:"opercode"`
	Text     string `json:"text"`
	Time     int64  `json:"time"`
	Worker   string `json:"worker"`
}
