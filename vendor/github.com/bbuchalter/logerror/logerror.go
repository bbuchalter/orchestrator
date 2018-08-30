package logerror

import (
	"errors"
	"fmt"
	"log/syslog"
	"os"
	"runtime/debug"
	"time"
)

var globalLogLevel = DEBUG
var printStackTrace = false

// syslogWriter is optional, and defaults to nil (disabled)
var syslogLevel = ERROR
var syslogWriter *syslog.Writer

// SetPrintStackTrace enables/disables dumping the stack upon error logging
func SetPrintStackTrace(shouldPrintStackTrace bool) {
	printStackTrace = shouldPrintStackTrace
}

// Define the log levels from 0 for FATAL, up to 6 for DEBUG
const (
	FATAL LogLevel = iota
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

// TimeFormat for log output
const TimeFormat = "2006-01-02 15:04:05"

// LogLevel indicates the severity of a log entry
type LogLevel int

func Criticale(err error) error {
	return logErrorEntry(CRITICAL, err)
}

func Errore(err error) error {
	return logErrorEntry(ERROR, err)
}

// Fatale emits a FATAL level entry and exists the program
func Fatale(err error) error {
	logErrorEntry(FATAL, err)
	os.Exit(1)
	return err
}

// Fatalf emits a FATAL level entry and exists the program
func Fatalf(message string, args ...interface{}) error {
	logFormattedEntry(FATAL, message, args...)
	os.Exit(1)
	return errors.New(logFormattedEntry(CRITICAL, message, args...))
}

func (level LogLevel) String() string {
	switch level {
	case FATAL:
		return "FATAL"
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	}
	return "unknown"
}

// logErrorEntry emits a log entry based on given error object
func logErrorEntry(logLevel LogLevel, err error) error {
	if err == nil {
		// No error
		return nil
	}
	entryString := fmt.Sprintf("%+v", err)
	logEntry(logLevel, entryString)
	if printStackTrace {
		debug.PrintStack()
	}
	return err
}

// logEntry emits a formatted log entry
func logEntry(logLevel LogLevel, message string, args ...interface{}) string {
	entryString := message
	for _, s := range args {
		entryString += fmt.Sprintf(" %s", s)
	}
	return logFormattedEntry(logLevel, entryString)
}

// logFormattedEntry nicely formats and emits a log entry
func logFormattedEntry(logLevel LogLevel, message string, args ...interface{}) string {
	if logLevel > globalLogLevel {
		return ""
	}
	msgArgs := fmt.Sprintf(message, args...)
	entryString := fmt.Sprintf("%s %s %s", time.Now().Format(TimeFormat), logLevel, msgArgs)
	fmt.Fprintln(os.Stderr, entryString)

	if syslogWriter != nil {
		go func() error {
			if logLevel > syslogLevel {
				return nil
			}
			switch logLevel {
			case FATAL:
				return syslogWriter.Emerg(msgArgs)
			case CRITICAL:
				return syslogWriter.Crit(msgArgs)
			case ERROR:
				return syslogWriter.Err(msgArgs)
			case WARNING:
				return syslogWriter.Warning(msgArgs)
			case NOTICE:
				return syslogWriter.Notice(msgArgs)
			case INFO:
				return syslogWriter.Info(msgArgs)
			case DEBUG:
				return syslogWriter.Debug(msgArgs)
			}
			return nil
		}()
	}
	return entryString
}
