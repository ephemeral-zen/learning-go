package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

func (l *Logger) logf(loglevel Level, format string, args ...any) {
	fullFormat := "[%s] " + format + "\n"
	fullArgs := append([]any{loglevel}, args...)
	_, _ = fmt.Fprintf(l.output, fullFormat+"\n", fullArgs...)
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

//func (l *Logger) Fatalf(format string, args ...any) {
//	if l.threshold > LevelFatal {
//		return
//	}
//	setDefaultOutput(l)
//	l.logf(format, args...)
//}

// Create a new instance of the Logger struct and return its address in memory
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}
	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}
