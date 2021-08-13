package loggy_test

import (
	"testing"

	"github.com/ImTomEddy/loggy"
)

func TestLogWithField(t *testing.T) {
	l := loggy.New(loggy.Opts{})
	checkLoggy(t, l)

	log := l.Default()
	checkLog(t, log)

	key := "key"
	value := "value"

	log.WithField(key, value)

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

func TestLogWithFields(t *testing.T) {
	l := loggy.New(loggy.Opts{})
	checkLoggy(t, l)

	log := l.Default()
	checkLog(t, log)

	key := "key"
	value := "value"

	log.WithFields(map[string]interface{}{
		key: value,
	})

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
