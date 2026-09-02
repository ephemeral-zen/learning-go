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

// (l* Logger) is a reciever here, we're basically defining a method for the Logger struct
func (l *Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(format, args...)
}

// Create a new instance of the Logger struct and return its address in memory
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}
