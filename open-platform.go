package weasn

import (
	"github.com/prodbox/weasn/kernel/context"
	"github.com/prodbox/weasn/official-account"
	"github.com/prodbox/weasn/open-platform"
	"github.com/prodbox/weasn/open-platform/base"
	"github.com/prodbox/weasn/open-platform/server"
)

// 微信开放平台
type OpenPlatform interface {
	// 获取用户授权页 URL
	GetPreAuthorizationUrl(redirectUri string) (string, error)
	// 使用授权码换取接口调用凭据和授权信息
	HandleAuthorize(authCode string) (*base.QueryAuth, error)
	// 获取授权方的帐号基本信息
	GetAuthorizer(authorizerAppId string) (*base.GetAuthorizerInfo, error)
	//获取授权方的选项设置信息
	GetAuthorizerOption(authorizerAppId, name string) (*base.GetAuthorizerOption, error)
	// 获取已授权的授权方列表
	GetAuthorizers(offset, count int) (*base.AuthorizerList, error)
	// 代公众号实现业务(单例)
	OfficialAccount(appId, refreshToken string) *official_account.Application
	// 服务端(单例)
	Server(opts ...server.Option) server.Server
	// AccessToken(单例)
	AccessToken() context.AccessToken
}

func NewOpenPlatform(opts ...context.Option) OpenPlatform {
	return open_platform.New(opts...)
}
