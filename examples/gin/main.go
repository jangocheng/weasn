package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prodbox/weasn"
	"github.com/prodbox/weasn/kernel/message"
	"github.com/prodbox/weasn/official-account/server"
)

func main() {
	app := weasn.NewOfficialAccount(
		weasn.AppId("wx71d408cd3aa9db9c"),
		weasn.Secret("9c918fe2434188a61ffaacc2f82d26d1"),
		weasn.Token("weiphp"),
		weasn.AESKey("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopq"),
	)

	serverOptions := []server.Option{
		server.MessageHandler(&messageHandler{}),
	}

	engine := gin.New()
	engine.GET("/wx/index", gin.WrapH(app.Server(serverOptions...)))
	engine.POST("/wx/index", gin.WrapH(app.Server(serverOptions...)))
	engine.Run(":8080")

	// qcYI2lIk2w34FUVRWq7_MdKUSE1VQxEAk_ZyrXtKB91Nabc1S3CrnBe44z6M57h2
	//fmt.Println(app.Base().GetValidIps())
}

type messageHandler struct{}

func (this *messageHandler) Text(mixed server.Mixed) message.Message {
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
