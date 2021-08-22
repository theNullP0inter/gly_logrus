package logrus

import (
	"testing"
)

func TestNewGooglyLogrusLogger(t *testing.T) {
	logrusLogger := NewGooglyLogrusLogger()

	logrusLogger.WithData(map[string]interface{}{"foo": "bar"}).Infof("test log")
}
