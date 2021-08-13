package main

import (
	"time"

	"github.com/ImTomEddy/loggy"
)

func main() {
	logger := loggy.New(loggy.Opts{
		LogLevel: loggy.LogLevelDebug,
		DefaultStaticFields: map[string]interface{}{
			"a.default.field": "default",
		},
		DefaultVariableFields: map[string]func() interface{}{
			"@timestamp": timestamp,
		},
	})

	logger.
		WithField("parent.child", "a value").
		WithField("parent.second", "another value").
		WithField("root", "a root field").
		Log(loggy.LogLevelDebug, "Hello World")
}

func timestamp() interface{} {
	return time.Now()
}
