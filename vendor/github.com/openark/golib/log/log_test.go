package log

import (
	"testing"
	"time"
)

func TestNotice(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Notice("this is a Notice message")
	expectedResult = "1974-05-19 01:02:03 NOTICE this is a Notice message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Notice("this is a Notice message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 NOTICE this is a Notice message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestNoticef(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Noticef("this is a Notice message")
	expectedResult = "1974-05-19 01:02:03 NOTICE this is a Notice message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Noticef("this is a Notice message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 NOTICE this is a Notice message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestDebug(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Debug("this is a Debug message")
	expectedResult = "1974-05-19 01:02:03 DEBUG this is a Debug message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Debug("this is a Debug message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 DEBUG this is a Debug message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestDebugf(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Debugf("this is a Debug message")
	expectedResult = "1974-05-19 01:02:03 DEBUG this is a Debug message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Debugf("this is a Debug message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 DEBUG this is a Debug message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestInfo(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Info("this is an Info message")
	expectedResult = "1974-05-19 01:02:03 INFO this is an Info message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Info("this is an Info message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 INFO this is an Info message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestInfof(t *testing.T) {
	stubNow()
	var result, expectedResult string

	// Test no args
	result = Infof("this is an Info message")
	expectedResult = "1974-05-19 01:02:03 INFO this is an Info message"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Infof("this is an Info message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 INFO this is an Info message with variables a b c"
	if result != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestWarning(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Warning("this is an Warning message")
	expectedResult = "1974-05-19 01:02:03 WARNING this is an Warning message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Warning("this is an Warning message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 WARNING this is an Warning message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestWarningf(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Warningf("this is an Warning message")
	expectedResult = "1974-05-19 01:02:03 WARNING this is an Warning message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Warningf("this is an Warning message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 WARNING this is an Warning message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestError(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Error("this is an Error message")
	expectedResult = "1974-05-19 01:02:03 ERROR this is an Error message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Error("this is an Error message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 ERROR this is an Error message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestErrorf(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Errorf("this is an Error message")
	expectedResult = "1974-05-19 01:02:03 ERROR this is an Error message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Errorf("this is an Error message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 ERROR this is an Error message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestCritical(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Critical("this is an Critical message")
	expectedResult = "1974-05-19 01:02:03 CRITICAL this is an Critical message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Critical("this is an Critical message with variables", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 CRITICAL this is an Critical message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func TestCriticalf(t *testing.T) {
	stubNow()
	var result error
	var expectedResult string

	// Test no args
	result = Criticalf("this is an Critical message")
	expectedResult = "1974-05-19 01:02:03 CRITICAL this is an Critical message"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}

	// Test with args
	result = Criticalf("this is an Critical message with variables %s %s %s", "a", "b", "c")
	expectedResult = "1974-05-19 01:02:03 CRITICAL this is an Critical message with variables a b c"
	if result.Error() != expectedResult {
		t.Errorf("Expected log output of '%s' but got '%s'", expectedResult, result)
	}
}

func stubNow() {
	// Overwrite the package level now variable with a fixed time
	now = func() time.Time {
		return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
	}
}
