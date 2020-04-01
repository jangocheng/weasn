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
		weasn.AppId("wx67426c1793b21554"),
		weasn.Secret("f2967a76bdcb9b9a4ec7d3c27794b384"),
		weasn.Token("maimi"),
		weasn.AESKey("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopq"),
	)

	serverOptions := []server.Option{
		server.MessageHandler(&messageHandler{}),
	}

	engine := gin.New()
	engine.GET("/mp/index", gin.WrapH(app.Server(serverOptions...)))
	engine.POST("/mp/index", gin.WrapH(app.Server(serverOptions...)))
	engine.Run(":8080")

	// qcYI2lIk2w34FUVRWq7_MdKUSE1VQxEAk_ZyrXtKB91Nabc1S3CrnBe44z6M57h2
	//fmt.Println(app.Base().GetValidIps())
}

type messageHandler struct{}

func (this *messageHandler) Text(mixed server.Mixed) message.Message {

	fmt.Println("Text = > ", app.CustomerService().Message(message.NewText("我是客服消息一")).To(mixed.FromUserName).Send())
	fmt.Println("Text = > ", app.CustomerService().Message(message.NewText("我是客服消息二")).To(mixed.FromUserName).Send())

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
