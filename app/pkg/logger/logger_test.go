package logger

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	got := SetLogger()
	got.Info("Logger Coverage")
	got.Error("Logger Coverage")
	if got == nil {
		t.Errorf("Got and Expected are not equals. Got: Log, expected: nil")
	}
}
