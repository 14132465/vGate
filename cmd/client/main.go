package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yz778899/vGate/net"
)

func main() {

	time.Sleep(time.Millisecond * 1000 * 3)
	start := time.Now()
	// 启动 1000 个客户端连续发送登录消息
	for i := 0; i < 10000; i++ {
		go func(id int) {
			app := net.NewAppClient().Config("ws://test-app:5566/")
			app.Handler(&ClientHandler{}).Connect(func(conn *websocket.Conn) {
				app.ConnSession.SendToGate("/user/login", LoginMsg())
			})
		}(i)

		if i%100 == 0 {
			fmt.Printf("已启动 %d 个 goroutine\n", i)
			//runtime.Gosched()
			time.Sleep(time.Millisecond * 1000)
		}
		time.Sleep(time.Millisecond * 1)
	}

	fmt.Printf("所有 goroutine 完成， 耗时 %v , \n", time.Since(start))

	for {
		time.Sleep(1 * time.Second * 5)
		//fmt.Println("等待中...")
	}
}
