package main

import (
	"github.com/14132465/vGate/net"
	_ "github.com/14132465/vGate/net/app"
	"github.com/14132465/vGate/net/handler"
)

func main() {

	handler := handler.GateHandler{}
	net.NewWsServer().Config(8080, "/").Handler(&handler).Run()

}
