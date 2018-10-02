package log

import (
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	// Overwrite the package level now variable with a fixed time
	now = func() time.Time {
		return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
	}

	result := Debug("this is a debug message")
	expectedResult := "1974-05-19 01:02:03 DEBUG this is a debug message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}
