# loggy
Loggy is a simple golang based structured logging package. It was built to take 
a json key using a `.` as a seperator to indicate a nested field.

## Example

```golang
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
```

```json
{
  "@timestamp": "2021-08-13T19:05:54.126701+01:00",
  "a": {
    "default": {
      "field": "default"
    }
  },
  "log": {
    "level": "DEBUG"
  },
  "message": "Hello World",
  "parent": {
    "child": "a value",
    "second": "another value"
  },
  "root": "a root field"
}
```

## License
Loggy is licensed under the [MIT](LICENSE) license.