package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct{}

func (this *Client) CreateClient() {
	// 连接服务器
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()

	// 发送消息
	go func() {
		for {
			err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				log.Println("发送失败:", err)
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// 接收消息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("接收失败:", err)
			return
		}
		fmt.Printf("收到回复: %s\n", msg)
	}
}
