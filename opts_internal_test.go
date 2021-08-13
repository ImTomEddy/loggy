package loggy

import (
	"fmt"
	"testing"
)

func TestOptsGetLogLevel(t *testing.T) {
	levels := map[LogLevel]LogLevel{
		0:             defaultLogLevel,
		LogLevelDebug: LogLevelDebug,
		LogLevelInfo:  LogLevelInfo,
		LogLevelWarn:  LogLevelWarn,
		LogLevelError: LogLevelError,
	}

	for in, out := range levels {
		t.Run(fmt.Sprintf("log level %v", in), func(t *testing.T) {
			opts := Opts{
				LogLevel: in,
			}

			optsOut := opts.getLogLevel()

			if out != optsOut {
				t.Logf("expected %v got %v", out, optsOut)
				t.Fail()
			}
		})
	}
}

func TestOptsGetLogLevelField(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		opts := Opts{}

		if defaultLogLevelField != opts.getLogLevelField() {
			t.Fail()
		}
	})
	t.Run("custom", func(t *testing.T) {
		field := "custom"
		opts := Opts{LogLevelField: field}

		if field != opts.getLogLevelField() {
			t.Fail()
		}
	})
}

func TestOptsGetDefaultStaticFields(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		opts := Opts{}

		if opts.getDefaultStaticFields() != nil {
			t.Fail()
		}
	})

	t.Run("not nil", func(t *testing.T) {
		opts := Opts{DefaultStaticFields: map[string]interface{}{}}

		if opts.getDefaultStaticFields() == nil {
			t.Fail()
		}
	})
}

func TestOptsGetDefaultVariableFields(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		opts := Opts{}

		if opts.getDefaultVariableFields() != nil {
			t.Fail()
		}
	})

	t.Run("not nil", func(t *testing.T) {
		opts := Opts{DefaultVariableFields: map[string](func() interface{}){}}

		if opts.getDefaultVariableFields() == nil {
			t.Fail()
		}
	})
}

func TestOptsGetMessageField(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		opts := Opts{}

		if defaultMessageField != opts.getMessageField() {
			t.Fail()
		}
	})
	t.Run("custom", func(t *testing.T) {
		field := "custom"
		opts := Opts{MessageField: field}

		if field != opts.getMessageField() {
			t.Fail()
		}
	})
}
