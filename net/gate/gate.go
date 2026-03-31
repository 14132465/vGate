package app

import "github.com/14132465/vGate/net/data"

//入口
type gate struct {
	//SessionManager *data.SessionManager
}

var (
	VGate          = &gate{}
	SessionManager *data.SessionManager
	//id    = 0
)

func init() {
	SessionManager = data.SessionManagerInstance
}
