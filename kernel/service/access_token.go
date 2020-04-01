package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/prodbox/weasn/kernel/context"
)

type accessToken struct {
	context.IAuthObject
}

func NewAccessToken(object context.IAuthObject) *accessToken {
	return &accessToken{IAuthObject: object}
}

// GetRefreshedToken 强制刷新
func (this *accessToken) GetRefreshedToken() (string, error) {
	return this.GetToken(true)
}

// GetToken 获取token
func (this *accessToken) GetToken(refresh ...bool) (string, error) {
	isRefresh := len(refresh) > 0 && refresh[0] == true

	// 获取缓存驱动
	cache := this.Options().Cache

	// 缓存获取
	if isRefresh == false && cache.IsExist(this.getCacheKey()) == true {
		if v := cache.Get(this.getCacheKey()); v != nil {
			return v.(string), nil
		}
	}

	// 微信获取
	var (
		formatted map[string]interface{}
		expiresIn int64 = 7200
		err       error
	)

	if formatted, err = this.request(); err != nil {
		return "", err
	}

	// 过期时间
	if v, ok := formatted["expires_in"]; ok {
		expiresIn = int64(v.(float64))
	}

	tokenStr := formatted[this.TokenKey()].(string)

	// 设置缓存
	if err := this.SetToken(tokenStr, expiresIn); err != nil {
		return "", fmt.Errorf("setToken fail: %s", err)
	}
	return tokenStr, nil
}

// SetToken 设置缓存
func (this *accessToken) SetToken(v string, lifetime int64) error {
	return this.Options().Cache.Set(this.getCacheKey(), v, lifetime-500)
}

// GetParams
func (this *accessToken) GetTokenParams() (map[string]string, error) {
	var (
		value string
		err   error
	)
	if value, err = this.GetToken(); err != nil {
		return nil, err
	}
	return map[string]string{this.QueryName(): value}, nil
}

func (this *accessToken) request() (map[string]interface{}, error) {
	response, err := this.Options().HttpClient.R().SetQueryParams(this.Credentials()).Execute(this.Method(), this.Endpoint())
	if err != nil {
		return nil, err
	}

	formatted := make(map[string]interface{})
	if err := json.Unmarshal(response.Body(), &formatted); err != nil {
		return nil, err
	}

	// 验证请求
	if _, ok := formatted[this.TokenKey()]; ok == false {
		return nil, fmt.Errorf("request access_token fail: %s", string(response.Body()))
	}

	return formatted, err
}

// getCacheKey 缓存键名
func (this *accessToken) getCacheKey() string {
	credentialsByte, _ := json.Marshal(this.Credentials())
	return fmt.Sprintf("weasn.open_platform.access_token.%s", this.md5(credentialsByte))
}

// md5 加密
func (this *accessToken) md5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
