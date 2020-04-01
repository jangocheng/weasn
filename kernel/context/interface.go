package context

import "net/http"

type Context interface {
	Options() Options

	Writer() http.ResponseWriter
	Request() *http.Request
	Reset(http.ResponseWriter, *http.Request)

	Query(key string) string

	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
	XML(code int, obj interface{})
	String(code int, format string, values ...interface{})
}

// 微信消息加密器
type Encrypter interface {
	// 加密
	Encrypt(content []byte) ([]byte, error)
	// 解密
	Decrypt(content, msgSignature, timestamp, nonce string) ([]byte, error)
}

type AccessToken interface {
	// GetRefreshedToken 强制刷新Token
	GetRefreshedToken() (string, error)
	// GetToken 获取token
	GetToken(...bool) (string, error)
	// SetToken 设置token缓存
	SetToken(string, int64) error
	// GetParams 获取token请求参数
	GetTokenParams() (map[string]string, error)
}

type IAuthObject interface {
	// Method 选项
	Options() Options
	// Method 请求类型 GET|POST
	Method() string
	// Key 返回Toekn数组键名
	TokenKey() string
	// QueryName
	QueryName() string
	// Endpoint Token请求地址
	Endpoint() string
	// Credentials 请求参数
	Credentials() map[string]string
}
