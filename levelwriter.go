package zerolog

/** BEGIN CUSTOM CODE */

import (
	"io"
)

// FilteredLevelWriter will only output messages that have a specific level.
//
// This writer can be used to filter specific message levels into a particular output file or stream. You can specify
// as many or as few levels as you'd like. Any messages with levels that are NOT a part of this list will NOT be
// logged regardless of what the logger's global level is set to.
//
// It should be noted that in order for a message to be processed by this writer, the logger itself must be set to
// the correct level as well. For example, if the logger is only set to log warning or greater messages, debug and
// info messages will still not be logged by this writer as they will not ever reach it. The general rule is to set
// the logger to the lowest level of messages you'd like to log and then filter higher level messages out as desired.
type FilteredLevelWriter struct {
	levels []Level
	writer io.Writer
}

// NewFilteredLevelwriter creates and returns a new FilteredLevelWriter object.
func NewFilteredLevelWriter(levels []Level, w io.Writer) *FilteredLevelWriter {
	return &FilteredLevelWriter{
		levels: levels,
		writer: w,
	}
}

// SetLevels updates the levels of messages that will be logged by this writer.
func (w *FilteredLevelWriter) SetLevels(levels []Level) {
	w.levels = levels
}

// Write simply writes the message out.
func (w *FilteredLevelWriter) Write(p []byte) (int, error) {
	return w.writer.Write(p)
}

// WriteLevel is called to write a message at a particular level.
//
// Only messages matching one of the filter levels will be logged. Other messages are effectively discarded.
func (w *FilteredLevelWriter) WriteLevel(level Level, p []byte) (int, error) {
	for _, l := range w.levels {
		if level == l {
			return w.Write(p)
		}
	}
	return len(p), nil
}

/** END CUSTOM CODE */
