package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/official-account/server"
)

var app weasn.OfficialAccount

func main() {

	app = weasn.NewOfficialAccount(
		weasn.AppId("123"),
		weasn.Secret("123"),
		weasn.Token("123"),
		weasn.AESKey("122"),
	)

	serverOptions := []server.Option{
		server.MessageHandler(&messageHandler{}),
	}

	engine := gin.New()
	engine.GET("/mp/index", gin.WrapH(app.Server(serverOptions...)))
	engine.POST("/mp/index", gin.WrapH(app.Server(serverOptions...)))
	engine.StaticFile("/MP_verify_EpghLbeN73ZTgbXh.txt", "examples/official-account/MP_verify_EpghLbeN73ZTgbXh.txt")

	engine.GET("/index", func(ctx *gin.Context) {
		// 开始授权
		app.OAuth().Redirect(ctx.Writer, ctx.Request, "http://chao.saaser.cn/oauth")
	})

	engine.GET("/oauth", func(ctx *gin.Context) {
		if user, err := app.OAuth().User(ctx.Query("code")); err == nil {
			fmt.Println(user)
			ctx.String(200, fmt.Sprintf("嗨！%s 您已授权成功", user.Nickname))
			return
		}
		ctx.String(200, "授权失败")
	})
	engine.Run(":8080")

	// qcYI2lIk2w34FUVRWq7_MdKUSE1VQxEAk_ZyrXtKB91Nabc1S3CrnBe44z6M57h2
	//fmt.Println(app.Base().GetValidIps())
}

type messageHandler struct{}

func (this *messageHandler) Text(mixed server.Mixed) message.Message {

	if mixed.Content == "模板消息" {
		fmt.Println(app.TemplateMessage().Send(map[string]interface{}{
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
		}))
	} else if mixed.Content == "客服消息" {
		fmt.Println("Text = > ", app.CustomerService().Message(message.NewText("我是客服消息一")).To(mixed.FromUserName).Send())
	} else if mixed.Content == "用户" {
		//fmt.Println(app.User().Get(mixed.FromUserName))
		fmt.Println(app.User().Select([]string{mixed.FromUserName}))
	}

	fmt.Println("openid => ", mixed.FromUserName)
	return message.NewText("我是文本消息 - > " + mixed.Content)
}

func (this *messageHandler) Image(mixed server.Mixed) message.Message {
	return message.NewImage(mixed.MediaId)
}

func (this *messageHandler) Voice(msg server.Mixed) message.Message {
	return message.NewText(fmt.Sprintf("您好！%s 太TM吓人了", msg.FromUserName))
}

func (this *messageHandler) Video(message server.Mixed) message.Message {
	return nil
}

func (this *messageHandler) ShortVideo(message server.Mixed) message.Message {
	return nil
}

func (this *messageHandler) Location(message server.Mixed) message.Message {
	return nil
}

func (this *messageHandler) Link(message server.Mixed) message.Message {
	return nil
}

func (this *messageHandler) Event(message server.Mixed) message.Message {
	return nil
}
