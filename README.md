# loggy
Loggy is a simple golang based structured logging package. It was built to take 
a json key using a `.` as a seperator to indicate a nested field.

## Example

```golang
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
```

```json
{
  "@timestamp": "2021-08-12T19:03:26.861871+01:00",
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