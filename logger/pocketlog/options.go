package pocketlog

import "io"

// Option is a custom function type used to configure a Logger.
// Here the Functional Options Pattern is used
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
