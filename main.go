package main

import (
	"github.com/14132465/vGate/net"
	"github.com/14132465/vGate/net/handler"
)

func main() {

	// 	type Message struct {
	// 		Type    string          `json:"type"`
	// 		Payload json.RawMessage `json:"payload"` // 保留原始 JSON
	// 	}

	// 	jsonData := `{
	//     "type": "publish",
	//     "payload": {
	//         "clientId": "client123",
	//         "topic": "test",
	//         "content": "hello world"
	//     }
	// }`

	// 	var msg Message
	// 	json.Unmarshal([]byte(jsonData), &msg)

	// 	fmt.Println(msg.Type)            // "publish"
	// 	fmt.Println(string(msg.Payload)) // {"clientId":"client123","topic":"test","content":"hello world"}

	handler := handler.GateHandler{}
	net.NewWsServer().Config(8080, "/").Handler(&handler).Run()

}
