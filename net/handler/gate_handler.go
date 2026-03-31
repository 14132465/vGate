package handler

import (
	"fmt"
	"net/http"

	"github.com/14132465/vGate/net/data"
	app "github.com/14132465/vGate/net/gate"
)

// GateHandler网关处理器，负责处理WebSocket连接和消息
type GateHandler struct {
}

// 收到消息
func (this *GateHandler) OnMessage(msg *data.WsMsg) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("处理消息时发生错误: %v\n", err)
		}
	}()

	switch msg.Cmd {
	case data.Subscription:
	//订阅消息
	case data.Publish:
		//发布消息
	case data.UnSubscription:
		//取消订阅消息
	case data.Notice:
		//通知消息
	case data.Request:
		//请求消息
	case data.Response:
		//回复消息
	default:
		//fmt.Printf("未知的消息指令 %v ", msg.Cmd)

	}

}

// 连接建立
func (this *GateHandler) OnConnect(w *http.ResponseWriter, r *http.Request) *data.Session {
	fmt.Printf("  main  ---- handler :  OnConnect  \n")
	// 将新连接添加到会话管理器

	session := app.SessionManager.AddSession(&data.Session{
		UUID:   -1,
		Status: 1,
		Resp:   w,
		Req:    r,
	})
	return session
}

// 连接断开
func (this *GateHandler) OnDisconnect(session *data.Session) {
	fmt.Printf("  main  ---- handler :  OnDisconnect  \n")
}
