package main

import (
	"encoding/json"
	"fmt"
	ws "test/net"
	"test/net/data"
)

func main() {

	type Message struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"` // 保留原始 JSON
	}

	jsonData := `{
    "type": "publish",
    "payload": {
        "clientId": "client123",
        "topic": "test",
        "content": "hello world"
    }
}`

	var msg Message
	json.Unmarshal([]byte(jsonData), &msg)

	fmt.Println(msg.Type)            // "publish"
	fmt.Println(string(msg.Payload)) // {"clientId":"client123","topic":"test","content":"hello world"}

	ws.NewWsServer().Config(8080, "/").Handler(func(msg data.WsMsg) {

		fmt.Printf("  main  ---- handler :  msg = %#v  \n", msg)

		// if nd, ok := msg.(data.NoDecoderMsg); ok {
		// 	data.Decoder(nd)
		// } else {
		// 	fmt.Printf("无法解析的消息类型 %v ", msg)
		// }

	}).Run()

}
