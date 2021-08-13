package loggy

type LogLevel int

const (
	LogLevelDebug LogLevel = 1
	LogLevelInfo  LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelError LogLevel = 4
)

func LogLevelToString(logLevel LogLevel) string {
	return map[LogLevel]string{
		LogLevelDebug: "DEBUG",
		LogLevelInfo:  "INFO",
		LogLevelWarn:  "WARN",
		LogLevelError: "ERROR",
	}[logLevel]
}

func StringToLogLevel(str string) LogLevel {
	return map[string]LogLevel{
		"DEBUG": LogLevelDebug,
		"INFO":  LogLevelInfo,
		"WARN":  LogLevelWarn,
		"ERROR": LogLevelError,
	}[str]
}
