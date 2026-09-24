package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold  Level
	output     io.Writer
	maxSymbols int
}

type LogEntry struct {
	LogLevel Level  `json:"level"`
	Message  string `json:"message"`
}

func (l *Logger) logf(loglevel Level, format string, args ...any) {
	//fullFormat := "%s " + format
	//fullArgs := append([]any{loglevel}, args...)
	fullMessage := fmt.Sprintf(format, args...)
	truncatedOutput := truncateLogOutput(fullMessage, l.maxSymbols)

	logEntry := LogEntry{
		LogLevel: loglevel,
		Message:  truncatedOutput,
	}

	logLine, _ := json.Marshal(logEntry)
	_, _ = fmt.Fprintln(l.output, string(logLine))
}

func truncateLogOutput(logOutput string, maxLen int) string {
	runes := []rune(logOutput)
	if len(runes) <= maxLen {
		return logOutput
	}
	return string(runes[:maxLen])
}

func setDefaultOutput(l *Logger) {
	if l.output == nil {
		l.output = os.Stdout
	}
}

// (l* Logger) is a reciever here, we're basically defining a method for the Logger struct
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	setDefaultOutput(l)
	l.logf(LevelDebug, format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	setDefaultOutput(l)
	l.logf(LevelInfo, format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}
	setDefaultOutput(l)
	l.logf(LevelWarn, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	setDefaultOutput(l)
	l.logf(LevelError, format, args...)
}

// Create a new instance of the Logger struct and return its address in memory
func New(threshold Level, opts ...Option) *Logger {
	// Define default values
	logger := &Logger{threshold: threshold, output: os.Stdout, maxSymbols: 1000}
	for _, configFunc := range opts {
		configFunc(logger)
	}
	return logger
}
