package pocketlog

// Level represents an available logging level
type Level byte

const (
	// enumeration
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)
