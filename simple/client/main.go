package main

import (
	"fmt"

	"github.com/14132465/vGate/net/data"
	"github.com/14132465/vGate/net/handler"
)

func main() {

	registry := handler.NewRegistry()

	registry.Register(handler.MsgHandlerCreate{
		Topic:      "/user/register",
		CreateFunc: handler.NewBaseMsgHandler,
	})

	topic := "/user/login"
	registry.Register(handler.MsgHandlerCreate{
		Topic:      topic,
		CreateFunc: handler.NewBaseMsgHandler,
	})

	creator, ok := registry.GetMsgHandlerCreate(topic)
	if ok {
		hdl := creator.CreateFunc(topic, &data.Session{}, &data.WsMsg{})

		fmt.Printf(" create a BaseMsgHandler %#v \n", hdl)
	}

}
