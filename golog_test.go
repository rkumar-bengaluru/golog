package golog_test

import (
	"golog"
	"testing"
)

func TestGoLogInfo(t *testing.T) {
	logger := golog.Default()
	logger.Info("Hello from logger")
}
