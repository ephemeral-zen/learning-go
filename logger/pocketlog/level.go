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

// Using Stringer interface here to turn Level to strings
func (l Level) String() string {
	//tp do: implement a logic to turn LevelY' int into a string
}
