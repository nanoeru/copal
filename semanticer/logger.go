package semanticer

import (
	"github.com/nanoeru/nlogger"
)

var stackLog = nlogger.NewLogger()

func init() {
	stackLog.AllIdle()
}

func Info(args ...interface{}) {
	nlogger.NewLogger().Info(args)
}

func Warn(args ...interface{}) {
	nlogger.NewLogger().Warn(args)
}

func Error(args ...interface{}) {
	nlogger.NewLogger().Error(args)
}

func Fatal(args ...interface{}) {
	nlogger.NewLogger().Fatal(args)
}
