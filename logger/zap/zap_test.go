package zap

import (
	"testing"

	"github.com/micro/go-micro/v2/logger"
)

func TestName(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	if l.String() != "zap" {
		t.Errorf("name is error %s", l.String())
	}

	t.Logf("test logger name: %s", l.String())
}

func TestLogf(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	l.Logf(logger.InfoLevel, "test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	l.SetLevel(logger.DebugLevel)
	l.Logf(logger.DebugLevel, "test show debug: %s", "debug msg")

	l.SetLevel(logger.InfoLevel)
	l.Logf(logger.DebugLevel, "test non-show debug: %s", "debug msg")
}