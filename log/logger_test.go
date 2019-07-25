package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel("error")
	DaoLogger.Info("dao")
	ServiceLogger.Warn("service")
	ControllerLogger.DebugF("controller")
}
