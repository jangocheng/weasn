# WeASN-阿萨辛

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)

微信开发SDK For Golang（公众号、小程序、微信支付、开放平台）仓库代码测试上传中!

| 目录 | 对应          | 功能                                                |
| ---- | ------------ | -------------------------------------------------- |
| official-account |微信公众平台| 网页授权、菜单、模板消息、消息回复、用户管理、消息转客服 |
| open-platform    |微信开放平台| 第三方平台推送事件、代公众号授权、代小程序授权|
| miniprogram      |微信小程序  | 授权、二维码|


## 获取

```sh
go get github.com/prodbox/weasn
```

## 微信公众号

### 快速接入

```go
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/official-account/server"
)

// 微信公众号实例
var officialAccount weasn.OfficialAccount

// 初始化公众号实例
officialAccount = weasn.NewOfficialAccount(
    weasn.AppId("123456789"),
    weasn.Secret("123456789"),
    weasn.Token("123456789"),
    weasn.AESKey("123456789"),
)

// gin接入
r := gin.New()
r.GET("/index", gin.WrapH(officialAccount.Server()))

```

## 微信开放平台

### 快速接入
```go
var openPlatform weasn.OpenPlatform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn"
	"github.com/prodbox/weasn/open-platform/server"
)

openPlatform = weasn.NewOpenPlatform(
    weasn.AppId("123456789"),
    weasn.Secret("123456789"),
    weasn.Token("123456789"),
    weasn.AESKey("123456789"),
)


s := app.Server().ComponentVerifyTicket(func(mixed server.Mixed) {
    fmt.Println("推送component_verify_ticket")
}).Authorized(func(mixed server.Mixed) {
  	fmt.Println("推送授权通知")
})
r := gin.New()
r.GET("/index", gin.WrapH(officialAccount.openPlatform()))

	
```