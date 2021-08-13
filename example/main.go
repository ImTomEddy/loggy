package main

import (
	"github.com/ImTomEddy/loggy"
)

func main() {
	logger := loggy.New(loggy.Opts{
		LogLevel:      loggy.LogLevelDebug,
		AutoTimestamp: true,
		DefaultFields: map[string]interface{}{
			"a.default.field": "default",
		},
	})

	logger.
		WithField("parent.child", "a value").
		WithField("parent.second", "another value").
		WithField("root", "a root field").
		Log(loggy.LogLevelDebug, "Hello World")
}
