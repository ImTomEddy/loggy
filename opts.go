package loggy

//Opts are the Loggy configuration options
type Opts struct {
	//LogLevel defines the minimum level to be logged by the logger
	LogLevel LogLevel
	//LogLevelField defines the field where the log level will be injected into
	LogLevelField string

	//DefaultStaticFields define fields which will be added to all log messages
	DefaultStaticFields map[string]interface{}
	//DefaultVariableFields define fields which will be injected to all log
	//messages with the function return values being the injected values
	DefaultVariableFields map[string](func() interface{})

	//MessageField defines the field where the log message will be injected into
	MessageField string
}

const (
	defaultLogLevel      = LogLevelInfo
	defaultMessageField  = "message"
	defaultLogLevelField = "log.level"
)

func (o *Opts) getLogLevel() LogLevel {
	if o.LogLevel == 0 {
		return defaultLogLevel
	}

	return o.LogLevel
}

func (o *Opts) getLogLevelField() string {
	if o.LogLevelField == "" {
		return defaultLogLevelField
	}

	return o.LogLevelField
}

func (o *Opts) getDefaultStaticFields() map[string]interface{} {
	return o.DefaultStaticFields
}

func (o *Opts) getDefaultVariableFields() map[string](func() interface{}) {
	return o.DefaultVariableFields
}

func (o *Opts) getMessageField() string {
	if o.MessageField == "" {
		return defaultMessageField
	}

	return o.MessageField
}
