package loggy

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Log struct {
	loggy  *Loggy
	fields map[string]interface{}
}

func (l *Log) WithField(key string, val interface{}) *Log {
	l.fields[key] = val
	return l
}

func (l *Log) WithFields(fields map[string]interface{}) *Log {
	for key, val := range fields {
		l.fields[key] = val
	}

	return l
}

func (l *Log) GetFields() map[string]interface{} {
	return l.fields
}

func (l *Log) Log(level LogLevel, message string, substitutes ...interface{}) {
	if level < l.loggy.Opts.getLogLevel() {
		return
	}

	if l.loggy.Opts.getDefaultStaticFields() != nil {
		for key, val := range l.loggy.Opts.getDefaultStaticFields() {
			l.fields[key] = val
		}
	}

	if l.loggy.Opts.getDefaultVariableFields() != nil {
		for key, fun := range l.loggy.Opts.getDefaultVariableFields() {
			l.fields[key] = fun()
		}
	}

	l.fields[l.loggy.Opts.getLogLevelField()] = LogLevelToString(l.loggy.Opts.getLogLevel())
	l.fields[l.loggy.Opts.getMessageField()] = fmt.Sprintf(message, substitutes...)
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
