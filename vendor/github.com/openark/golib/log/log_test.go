package log

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
	"time"
)

type loggingTestCase struct {
	description          string
	preTest              func(t *testing.T)
	expectedReturnValue  string
	expectedLoggedOutput string
	testSubject          func() string
}

var loggingTestCases = []*loggingTestCase{
	&loggingTestCase{
		description: "Notice logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 NOTICE this is a Notice message",
		expectedLoggedOutput: "1974-05-19 01:02:03 NOTICE this is a Notice message\n",
		testSubject:          func() string { return Notice("this is a Notice message") },
	},
	&loggingTestCase{
		description: "Notice logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 NOTICE this is a Notice message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 NOTICE this is a Notice message a b c\n",
		testSubject:          func() string { return Notice("this is a Notice message", "a", "b", "c") },
	},
	&loggingTestCase{
		description: "Noticef logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 NOTICE this is a Noticef message",
		expectedLoggedOutput: "1974-05-19 01:02:03 NOTICE this is a Noticef message\n",
		testSubject:          func() string { return Noticef("this is a Noticef message") },
	},
	&loggingTestCase{
		description: "Noticef logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 NOTICE this is a Noticef message",
		expectedLoggedOutput: "1974-05-19 01:02:03 NOTICE this is a Noticef message\n",
		testSubject:          func() string { return Noticef("%s %s a Noticef message", "this", "is") },
	},
	&loggingTestCase{
		description: "Debug logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 DEBUG this is a Debug message",
		expectedLoggedOutput: "1974-05-19 01:02:03 DEBUG this is a Debug message\n",
		testSubject:          func() string { return Debug("this is a Debug message") },
	},
	&loggingTestCase{
		description: "Debug logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 DEBUG this is a Debug message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 DEBUG this is a Debug message a b c\n",
		testSubject:          func() string { return Debug("this is a Debug message", "a", "b", "c") },
	},
	&loggingTestCase{
		description: "Debugf logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 DEBUG this is a Debugf message",
		expectedLoggedOutput: "1974-05-19 01:02:03 DEBUG this is a Debugf message\n",
		testSubject:          func() string { return Debugf("this is a Debugf message") },
	},
	&loggingTestCase{
		description: "Debugf logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 DEBUG this is a Debugf message",
		expectedLoggedOutput: "1974-05-19 01:02:03 DEBUG this is a Debugf message\n",
		testSubject:          func() string { return Debugf("%s %s a Debugf message", "this", "is") },
	},
	&loggingTestCase{
		description: "Info logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 INFO this is a Info message",
		expectedLoggedOutput: "1974-05-19 01:02:03 INFO this is a Info message\n",
		testSubject:          func() string { return Info("this is a Info message") },
	},
	&loggingTestCase{
		description: "Info logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 INFO this is a Info message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 INFO this is a Info message a b c\n",
		testSubject:          func() string { return Info("this is a Info message", "a", "b", "c") },
	},
	&loggingTestCase{
		description: "Infof logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 INFO this is a Infof message",
		expectedLoggedOutput: "1974-05-19 01:02:03 INFO this is a Infof message\n",
		testSubject:          func() string { return Infof("this is a Infof message") },
	},
	&loggingTestCase{
		description: "Infof logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 INFO this is a Infof message",
		expectedLoggedOutput: "1974-05-19 01:02:03 INFO this is a Infof message\n",
		testSubject:          func() string { return Infof("%s %s a Infof message", "this", "is") },
	},
	&loggingTestCase{
		description: "Warning logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 WARNING this is a Warning message",
		expectedLoggedOutput: "1974-05-19 01:02:03 WARNING this is a Warning message\n",
		testSubject:          func() string { return Warning("this is a Warning message").Error() },
	},
	&loggingTestCase{
		description: "Warning logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 WARNING this is a Warning message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 WARNING this is a Warning message a b c\n",
		testSubject:          func() string { return Warning("this is a Warning message", "a", "b", "c").Error() },
	},
	&loggingTestCase{
		description: "Warningf logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 WARNING this is a Warning message",
		expectedLoggedOutput: "1974-05-19 01:02:03 WARNING this is a Warning message\n",
		testSubject:          func() string { return Warningf("this is a Warning message").Error() },
	},
	&loggingTestCase{
		description: "Warningf logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 WARNING this is a Warning message",
		expectedLoggedOutput: "1974-05-19 01:02:03 WARNING this is a Warning message\n",
		testSubject:          func() string { return Warningf("%s %s a Warning message", "this", "is").Error() },
	},
	&loggingTestCase{
		description: "Error logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 ERROR this is a Error message",
		expectedLoggedOutput: "1974-05-19 01:02:03 ERROR this is a Error message\n",
		testSubject:          func() string { return Error("this is a Error message").Error() },
	},
	&loggingTestCase{
		description: "Error logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 ERROR this is a Error message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 ERROR this is a Error message a b c\n",
		testSubject:          func() string { return Error("this is a Error message", "a", "b", "c").Error() },
	},
	&loggingTestCase{
		description: "Errore logs errors",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "this is a Error message",
		expectedLoggedOutput: "1974-05-19 01:02:03 ERROR this is a Error message\n",
		testSubject:          func() string { return Errore(errors.New("this is a Error message")).Error() },
	},
	&loggingTestCase{
		description: "Errorf logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 ERROR this is a Error message",
		expectedLoggedOutput: "1974-05-19 01:02:03 ERROR this is a Error message\n",
		testSubject:          func() string { return Errorf("this is a Error message").Error() },
	},
	&loggingTestCase{
		description: "Errorf logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 ERROR this is a Error message",
		expectedLoggedOutput: "1974-05-19 01:02:03 ERROR this is a Error message\n",
		testSubject:          func() string { return Errorf("%s %s a Error message", "this", "is").Error() },
	},
	&loggingTestCase{
		description: "Critical logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 CRITICAL this is a Critical message",
		expectedLoggedOutput: "1974-05-19 01:02:03 CRITICAL this is a Critical message\n",
		testSubject:          func() string { return Critical("this is a Critical message").Error() },
	},
	&loggingTestCase{
		description: "Critical logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 CRITICAL this is a Critical message a b c",
		expectedLoggedOutput: "1974-05-19 01:02:03 CRITICAL this is a Critical message a b c\n",
		testSubject:          func() string { return Critical("this is a Critical message", "a", "b", "c").Error() },
	},
	&loggingTestCase{
		description: "Criticale logs errors",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "this is a Critical message",
		expectedLoggedOutput: "1974-05-19 01:02:03 CRITICAL this is a Critical message\n",
		testSubject:          func() string { return Criticale(errors.New("this is a Critical message")).Error() },
	},
	&loggingTestCase{
		description: "Criticalf logs messages",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 CRITICAL this is a Critical message",
		expectedLoggedOutput: "1974-05-19 01:02:03 CRITICAL this is a Critical message\n",
		testSubject:          func() string { return Criticalf("this is a Critical message").Error() },
	},
	&loggingTestCase{
		description: "Criticalf logs messages and additional variables",
		preTest: func(t *testing.T) {
			stubNow()
		},
		expectedReturnValue:  "1974-05-19 01:02:03 CRITICAL this is a Critical message",
		expectedLoggedOutput: "1974-05-19 01:02:03 CRITICAL this is a Critical message\n",
		testSubject:          func() string { return Criticalf("%s %s a Critical message", "this", "is").Error() },
	},
}

func TestLoggingTestCases(t *testing.T) {
	for _, testCase := range loggingTestCases {
		t.Run(testCase.description, func(t *testing.T) {

			// Run any code needed before the test begins
			if testCase.preTest != nil {
				testCase.preTest(t)
			}

			returnValue, loggedOutput := getLoggedOutput(t, testCase.testSubject)

			if returnValue != testCase.expectedReturnValue {
				t.Errorf("Expected return of '%s' but got '%s'", testCase.expectedReturnValue, returnValue)
			}
			if loggedOutput != testCase.expectedLoggedOutput {
				t.Errorf("Expected log output of '%s' but got '%s'", testCase.expectedLoggedOutput, loggedOutput)
			}
		})
	}
}
func TestLogLevelFromString(t *testing.T) {
	var result LogLevel
	var err error

	result, err = LogLevelFromString("FATAL")
	if err != nil {
		t.Error(err)
	}
	if result != 0 {
		t.Errorf("Expected LogLevel for FATAL to be 0, but was %s", result)
	}

	result, err = LogLevelFromString("CRITICAL")
	if err != nil {
		t.Error(err)
	}
	if result != 1 {
		t.Errorf("Expected LogLevel for CRITICAL to be 1, but was %s", result)
	}

	result, err = LogLevelFromString("ERROR")
	if err != nil {
		t.Error(err)
	}
	if result != 2 {
		t.Errorf("Expected LogLevel for ERROR to be 2, but was %s", result)
	}

	result, err = LogLevelFromString("WARNING")
	if err != nil {
		t.Error(err)
	}
	if result != 3 {
		t.Errorf("Expected LogLevel for WARNING to be 3, but was %s", result)
	}

	result, err = LogLevelFromString("NOTICE")
	if err != nil {
		t.Error(err)
	}
	if result != 4 {
		t.Errorf("Expected LogLevel for NOTICE to be 4, but was %s", result)
	}

	result, err = LogLevelFromString("INFO")
	if err != nil {
		t.Error(err)
	}
	if result != 5 {
		t.Errorf("Expected LogLevel for INFO to be 5, but was %s", result)
	}

	result, err = LogLevelFromString("DEBUG")
	if err != nil {
		t.Error(err)
	}
	if result != 6 {
		t.Errorf("Expected LogLevel for DEBUG to be 6, but was %s", result)
	}

	result, err = LogLevelFromString("INVALID")
	if err == nil {
		t.Error("Expected an error not none was returned.")
	}
	if err.Error() != "Unknown LogLevel name: INVALID" {
		t.Errorf("Expected 'Unknown LogLevel name: INVALID' but was '%s'", err.Error())
	}
	if result != 0 {
		t.Errorf("Expected a return value of 0, but was %s", result)
	}
}

func stubNow() {
	// Overwrite the package level now variable with a fixed time
	SetNow(func() time.Time {
		return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
	})
}

func getLoggedOutput(t *testing.T, testSubject func() string) (testSubjectReturnValue string, loggedOutput string) {
	// Keep a copy of the original logDestination
	defaultLogDestination := GetLogDestination()
	// Restore original logDestination when func completes
	defer SetLogDestination(defaultLogDestination)

	// Create a pipe to capture log output
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	SetLogDestination(writer)

	// copy the output in a separate goroutine so printing can't block indefinitely
	outputChannel := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		outputChannel <- buf.String()
	}()

	// Test subject
	testSubjectReturnValue = testSubject()

	// Close pipe and collect log output
	writer.Close()
	loggedOutput = <-outputChannel

	return
}
