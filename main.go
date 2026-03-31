package main

import (
	"github.com/14132465/vGate/net"
	"github.com/14132465/vGate/net/handler"
)

func main() {

	handler := handler.GateHandler{}
	net.NewWsServer().Config(8080, "/").Handler(&handler).Run()

}
