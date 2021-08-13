package loggy_test

import (
	"testing"

	"github.com/ImTomEddy/loggy"
)

func TestNew(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		l := loggy.New(loggy.Opts{})
		checkLoggy(t, l)
	})

	t.Run("custom options", func(t *testing.T) {
		l := loggy.New(loggy.Opts{
			TimeStampField: "time",
		})
		checkLoggy(t, l)
	})

}

func TestLoggyWithField(t *testing.T) {
	l := loggy.New(loggy.Opts{})
	checkLoggy(t, l)

	key := "key"
	value := "value"

	log := l.WithField(key, value)
	checkLog(t, log)

	if len(log.GetFields()) != 1 {
		t.Logf("expected 1 field got %v", len(log.GetFields()))
		t.Fail()
	}

	val, ok := log.GetFields()[key]

	if !ok {
		t.Logf("unable to find key %s", key)
		t.Fail()
	}

	if value != val {
		t.Logf("expected %s got %s", value, val)
		t.Fail()
	}
}

func TestLoggyWithFields(t *testing.T) {
	l := loggy.New(loggy.Opts{})
	checkLoggy(t, l)

	key := "key"
	value := "value"

	log := l.WithFields(map[string]interface{}{
		key: value,
	})
	checkLog(t, log)

	if len(log.GetFields()) != 1 {
		t.Logf("expected 1 field got %v", len(log.GetFields()))
		t.Fail()
	}

	val, ok := log.GetFields()[key]

	if !ok {
		t.Logf("unable to find key %s", key)
		t.Fail()
	}

	if value != val {
		t.Logf("expected %s got %s", value, val)
		t.Fail()
	}
}

func TestLoggyDefault(t *testing.T) {
	l := loggy.New(loggy.Opts{})
	checkLoggy(t, l)

	log := l.Default()
	checkLog(t, log)

	if len(log.GetFields()) != 0 {
		t.Logf("expected 0 fields got %v", len(log.GetFields()))
		t.Fail()
	}
}
