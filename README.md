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
    weasn.Cache(&customCache{}),      // 自定义缓存（默认文件缓存）
)

// 配合gin
r := gin.New()
r.GET("/index", gin.WrapH(officialAccount.Server()))

```

### 事件处理
```go

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn/kernel/message"
)

//
s := officialAccount.Server(
    server.MessageHandler(&messageHandler{}),
)

/*
type IMessage interface {
	Text(Mixed) message.Message
	Image(Mixed) message.Message
	Voice(Mixed) message.Message
	Video(Mixed) message.Message
	Link(Mixed) message.Message
	Event(Mixed) message.Message
	ShortVideo(Mixed) message.Message
	Location(Mixed) message.Message
}
*/

type messageHandler struct{}

func (this *messageHandler) Text(mixed server.Mixed) message.Message {
	return message.NewText("回复一条文本消息")
}

func (this *messageHandler) Image(mixed server.Mixed) message.Message {
	// 发送一条客服消息
	officialAccount.CustomerService().Message(message.NewText("我是客服消息")).To(mixed.FromUserName).Send()
	// 不回复内容
	return nil
}

func (this *messageHandler) Voice(msg server.Mixed) message.Message {
	return message.NewText("文本消息")
}
...

```

### 模板消息
```go
officialAccount.TemplateMessage().Send(map[string]interface{}{
			"touser":      mixed.FromUserName,
			"template_id": "3yInRq35ahx-VswdudnNfwiM-ncYn2myFnZS9JfVbs8",
			"data": map[string]interface{}{
				"first": map[string]string{
					"value": "恭喜中奖",
					"color": "#173177",
				},
				"keyword1": map[string]string{
					"value": "非洲一日游",
					"color": "#173177",
				},
				"keyword2": map[string]string{
					"value": "2017年9月8日",
					"color": "#173177",
				},
				"remark": map[string]string{
					"value": "请于3日内前往非洲",
					"color": "#173177",
				},
			},
		})

```

## 微信开放平台

### 快速接入
```go

// 微信开放平台实例
var openPlatform weasn.OpenPlatform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn"
	"github.com/prodbox/weasn/open-platform/server"
)

// 初始化开放平台实例
openPlatform = weasn.NewOpenPlatform(
    weasn.AppId("123456789"),
    weasn.Secret("123456789"),
    weasn.Token("123456789"),
    weasn.AESKey("123456789"),
)

// 添加监听事件
s := app.Server().ComponentVerifyTicket(func(mixed server.Mixed) {
    fmt.Println("推送component_verify_ticket")
}).Authorized(func(mixed server.Mixed) {
  	fmt.Println("推送授权通知")
})

// 配合gin
r := gin.New()
r.GET("/index", gin.WrapH(s))

```