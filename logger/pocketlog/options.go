package pocketlog

import (
	"io"
)

// Option is a custom function type used to configure a Logger.
// Here the Functional Options Pattern is used
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(logger *Logger) {
		logger.output = output
	}
}

func WithMaxSymbols(maxSymbTrunc int) Option {
	return func(logger *Logger) {
		logger.maxSymbols = maxSymbTrunc
	}
}
