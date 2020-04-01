package message

// 错误响应接口
type JsonError interface {
	GetCode() int64
	GetMsg() string
}

// 错误消息固定格式
type JsonErrorMessage struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetCode 错误码
func (e *JsonErrorMessage) GetCode() int64 {
	return e.ErrCode
}

// GetMsg 错误消息
func (e *JsonErrorMessage) GetMsg() string {
	return e.ErrMsg
}
