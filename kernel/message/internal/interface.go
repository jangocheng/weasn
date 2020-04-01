package internal

// 响应消息接口
type Response interface {
	To(to string)
	From(from string)
	MessageId(msgId int64)
	CreatedAt(int64)
}

// 客服消息接口
type Customer interface {
	To(to string)
	KfAccount(kf string)
}
