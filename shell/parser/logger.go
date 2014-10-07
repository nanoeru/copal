package parser

import (
	"github.com/nanoeru/nlogger"
)

var lexLog = nlogger.NewLogger()
var exprLog = nlogger.NewLogger()

func init() {
	lexLog.Info.Idle()
	exprLog.Info.Idle()
	//lexLog.AllIdle()
	//exprLog.AllIdle()
}
