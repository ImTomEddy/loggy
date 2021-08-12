package main

import (
	"github.com/ImTomEddy/loggy"
)

func main() {
	logger := loggy.NewLoggy(loggy.Options{
		LogLevel:      loggy.LogLevelDebug,
		AutoTimestamp: true,
		DefaultFields: loggy.Fields{
			"a.default.field": "default",
		},
	})

	logger.
		WithField("parent.child", "a value").
		WithField("parent.second", "another value").
		WithField("root", "a root field").
		Log(loggy.LogLevelDebug, "Hello World")
}
