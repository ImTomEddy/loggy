package loggy

type Loggy struct {
	Opts
}

func New(opts Opts) *Loggy {
	return &Loggy{Opts: opts}
}

func (l *Loggy) WithField(key string, val interface{}) *Log {
	return l.WithFields(map[string]interface{}{key: val})
}

func (l *Loggy) WithFields(fields map[string]interface{}) *Log {
	return &Log{
		loggy:  l,
		fields: fields,
	}
}

func (l *Loggy) Default() *Log {
	return &Log{
		loggy:  l,
		fields: map[string]interface{}{},
	}
}
