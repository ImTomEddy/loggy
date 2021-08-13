package loggy

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Log struct {
	loggy  *Loggy
	fields Fields
}

func (l *Log) WithField(key string, val interface{}) *Log {
	l.fields[key] = val
	return l
}

func (l *Log) WithFields(fields Fields) *Log {
	for key, val := range fields {
		l.fields[key] = val
	}

	return l
}

func (l *Log) GetFields() Fields {
	return l.fields
}

func (l *Log) Log(level LogLevel, message string, substitutes ...interface{}) {
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

	l.fields[l.loggy.logLevelField] = LogLevelToString(l.loggy.logLevel)
	l.fields[l.loggy.messageField] = fmt.Sprintf(message, substitutes...)
	fmt.Println(l.buildJSONLog())
}

func (l *Log) buildJSONLog() string {
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
