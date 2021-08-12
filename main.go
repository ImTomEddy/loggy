package loggy

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
)

var (
	logLevelMap = map[LogLevel]string{
		LogLevelDebug: "DEBUG",
		LogLevelInfo:  "INFO",
		LogLevelWarn:  "WARN",
		LogLevelError: "ERROR",
	}
)

type Options struct {
	LogLevel       LogLevel
	DefaultFields  Fields
	MessageField   string
	LogLevelField  string
	AutoTimestamp  bool
	TimeStampField string
}

type Loggy struct {
	logLevel       LogLevel
	defaultFields  Fields
	messageField   string
	logLevelField  string
	autoTimestamp  bool
	timestampField string
}

type log struct {
	loggy  *Loggy
	fields Fields
}

type Fields map[string]interface{}

func NewLoggy(opts Options) *Loggy {
	msgField := opts.MessageField

	if msgField == "" {
		msgField = "message"
	}

	return &Loggy{
		logLevel:       opts.LogLevel,
		defaultFields:  opts.DefaultFields,
		autoTimestamp:  opts.AutoTimestamp,
		timestampField: optOrDefaultStr(opts.TimeStampField, "@timestamp"),
		logLevelField:  optOrDefaultStr(opts.LogLevelField, "log.level"),
		messageField:   optOrDefaultStr(opts.MessageField, "message"),
	}
}

func optOrDefaultStr(optVal string, defaultVal string) string {
	if optVal == "" {
		return defaultVal
	}

	return optVal
}

func (l *Loggy) WithField(key string, val interface{}) *log {
	return l.WithFields(Fields{key: val})
}

func (l *Loggy) WithFields(fields Fields) *log {
	return &log{
		loggy:  l,
		fields: fields,
	}
}

func (l *log) WithField(key string, val interface{}) *log {
	l.fields[key] = val
	return l
}

func (l *log) WithFields(fields Fields) *log {
	for key, val := range fields {
		l.fields[key] = val
	}

	return l
}

func (l *log) Log(level LogLevel, message string, substitutes ...interface{}) {
	if level < l.loggy.logLevel {
		return
	}

	if l.loggy.defaultFields != nil {
		for key, val := range l.loggy.defaultFields {
			l.fields[key] = val
		}
	}

	if l.loggy.autoTimestamp {
		l.fields[l.loggy.timestampField] = time.Now()
	}

	l.fields[l.loggy.logLevelField] = logLevelMap[l.loggy.logLevel]
	l.fields[l.loggy.messageField] = fmt.Sprintf(message, substitutes...)
	fmt.Println(l.buildJSONLog())
}

func (l *log) buildJSONLog() string {
	nested := map[string]interface{}{}

	for key, val := range l.fields {
		setChild(nested, key, val)
	}

	b, err := json.Marshal(nested)

	if err != nil {
		return ""
	}

	return string(b)
}

func setChild(parent map[string]interface{}, key string, val interface{}) {
	keys := strings.Split(key, ".")

	if len(keys) == 0 {
		return
	}

	if len(keys) == 1 {
		parent[keys[0]] = val
		return
	}

	if _, ok := parent[keys[0]]; !ok {
		parent[keys[0]] = map[string]interface{}{}
	}

	switch child := parent[keys[0]].(type) {
	case map[string]interface{}:
		setChild(child, strings.ReplaceAll(key, fmt.Sprintf("%s.", keys[0]), ""), val)
	}
}
