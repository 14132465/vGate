package main

import (
	"time"

	"github.com/14132465/vGate/net/app"
	_ "github.com/14132465/vGate/net/app"
)

func main() {

	app.Log.Info(" first log  ")

	for {
		time.Sleep(time.Microsecond)
	}

}
