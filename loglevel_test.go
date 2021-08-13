package loggy_test

import (
	"strconv"
	"testing"

	"github.com/ImTomEddy/loggy"
)

func TestLogLevelToString(t *testing.T) {
	logLevels := map[loggy.LogLevel]string{
		loggy.LogLevelDebug: "DEBUG",
		loggy.LogLevelInfo:  "INFO",
		loggy.LogLevelWarn:  "WARN",
		loggy.LogLevelError: "ERROR",
		0:                   "",
	}

	for level, val := range logLevels {
		t.Run("log level "+strconv.Itoa(int(level)), func(t *testing.T) {
			str := loggy.LogLevelToString(level)

			if val != str {
				t.Logf("expected '%s' got '%s'", val, str)
				t.Fail()
			}
		})
	}
}

func TestStringToLoveLevel(t *testing.T) {
	logLevels := map[string]loggy.LogLevel{
		"DEBUG": loggy.LogLevelDebug,
		"INFO":  loggy.LogLevelInfo,
		"WARN":  loggy.LogLevelWarn,
		"ERROR": loggy.LogLevelError,
		"":      0,
	}

	for val, level := range logLevels {
		t.Run("string "+val, func(t *testing.T) {
			l := loggy.StringToLogLevel(val)

			if level != l {
				t.Logf("expected '%v' got '%v'", level, l)
				t.Fail()
			}
		})
	}
}
