package loggy_test

import (
	"testing"

	"github.com/ImTomEddy/loggy"
)

func checkLoggy(t *testing.T, l *loggy.Loggy) {
	if l == nil {
		t.Log("loggy is nil")
		t.Fail()
	}
}

func checkLog(t *testing.T, l *loggy.Log) {
	if l == nil {
		t.Log("log is nil")
		t.Fail()
	}
}
