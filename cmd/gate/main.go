package main

import (
	"fmt"

	"github.com/yz778899/vGate/net"
	"github.com/yz778899/vGate/net/env"
)

func main() {

	defer env.Log.Sync()

	err := net.NewWsServer().Run()

	//err := net.NewWsServer().Run()  //启用上面默认参数启动
	if err != nil {
		fmt.Printf("gate failed to start: %v ", err)
	}

}
