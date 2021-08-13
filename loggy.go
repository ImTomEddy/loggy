package loggy

type Opts struct {
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

type Fields map[string]interface{}

func New(opts Opts) *Loggy {
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

func (l *Loggy) WithField(key string, val interface{}) *Log {
	return l.WithFields(Fields{key: val})
}

func (l *Loggy) WithFields(fields Fields) *Log {
	return &Log{
		loggy:  l,
		fields: fields,
	}
}

func (l *Loggy) Default() *Log {
	return &Log{
		loggy:  l,
		fields: Fields{},
	}
}
