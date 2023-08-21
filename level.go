package slogwatermill

// A Level is the importance or severity of a log event.
// The higher the level, the more important or severe the event.
type Level int

const (
	LevelTrace Level = -8
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)
