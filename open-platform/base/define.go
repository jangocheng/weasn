package base

import "github.com/prodbox/weasn/kernel/message"

type FuncInfo struct {
	FuncscopeCategory struct {
		Id int64 `json:"id"`
	} `json:"funcscope_category"`
}

type TypeInfo struct {
	Id int64 `json:"id"`
}

type Categories struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

type Network struct {
	RequestDomain   []string `json:"RequestDomain"`
	WsRequestDomain []string `json:"WsRequestDomain"`
	UploadDomain    []string `json:"UploadDomain"`
	DownloadDomain  []string `json:"DownloadDomain"`
}

type MiniProgramInfo struct {
	Network     Network       `json:"network"`
	Categories  []*Categories `json:"categories"`
	VisitStatus int32         `json:"visit_status"`
}

type AuthorizationInfo struct {
	AuthorizerAppId        string      `json:"authorizer_appid"`
	AuthorizerAccessToken  string      `json:"authorizer_access_token"`
	ExpiresIn              float64     `json:"expires_in"`
	AuthorizerRefreshToken string      `json:"authorizer_refresh_token"`
	FuncInfo               []*FuncInfo `json:"func_info"`
}

type AuthorizerInfo struct {
	Nickname        string   `json:"nick_name"`
	HeadImg         string   `json:"head_img"`
	ServiceTypeInfo TypeInfo `json:"service_type_info"`
	VerifyTypeInfo  TypeInfo `json:"verify_type_info"`
	Username        string   `json:"user_name"`
	PrincipalName   string   `json:"principal_name"`
	BusinessInfo    struct {
		OpenStore int64 `json:"open_store"`
		OpenScan  int64 `json:"open_scan"`
		OpenPay   int64 `json:"open_pay"`
		OpenCard  int64 `json:"open_card"`
		OpenShake int64 `json:"open_shake"`
	} `json:"business_info"`
	Signature       string          `json:"signature"`
	MiniProgramInfo MiniProgramInfo `json:"MiniProgramInfo"`
	Alias           string          `json:"alias"`
	QRCodeUrl       string          `json:"qrcode_url"`
}

// 使用授权码换取公众号的授权信息
type QueryAuth struct {
	message.JsonErrorMessage
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

// 获取授权方信息
type GetAuthorizerInfo struct {
	message.JsonErrorMessage
	AuthorizerInfo    AuthorizerInfo `json:"authorizer_info"`
	AuthorizationInfo struct {
		AuthorizerAppId string      `json:"authorizer_appid"`
		FuncInfo        []*FuncInfo `json:"func_info"`
	} `json:"authorization_info"`
}

func (this *GetAuthorizerInfo) FuncInfoToSlice() []int64 {
	var catSlice []int64
	for _, v := range this.AuthorizationInfo.FuncInfo {
		catSlice = append(catSlice, v.FuncscopeCategory.Id)
	}
	return catSlice
}

// 获取授权方的选项设置信息
type GetAuthorizerOption struct {
	message.JsonErrorMessage
	AuthorizerAppId string `json:"authorizer_appid"`
	OptionName      string `json:"option_name"`
	OptionValue     string `json:"option_value"`
}

// 获取预授权码
type PreAuthCode struct {
	message.JsonErrorMessage
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

type AuthorizerList struct {
	message.JsonErrorMessage
	TotalCount int64 `json:"total_count"`
	List       []struct {
		AuthorizerAppId string `json:"authorizer_appid"`
		RefreshToken    string `json:"refresh_token"`
		AuthTime        string `json:"auth_time"`
	} `json:"list"`
}
