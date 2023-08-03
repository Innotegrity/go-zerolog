package zerolog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

/** BEGIN CUSTOM CODE */

// JSONWriter is a custom JSON-formatted writer that changes the default field names for this and only this
// writer.
//
// Unlike using the global `zerolog.CallerFieldName` value to change the caller name for all log writers, you
// can use the writer's `CallerFieldName` to change the name only for this writer.
//
// By default if a field name is not provided, the global field name will be used instead. When excluding fields
// from the output, be sure to use the customized field name if one was set.
//
// NOTE: Be aware that this writer may not be as efficient as the built-in JSON writer as the normal JSON is
// unmarshaled to a generic map[string]interface{} object and then re-marshaled after being updated.
type JSONWriter struct {
	// Out is the output destination.
	Out io.Writer

	// TimeFormat specifies the format for timestamp in output.
	TimeFormat string

	// ExcludeFields defines fields to exclude in the output.
	ExcludeFields []string

	// CallerFieldName defines the name of the caller field.
	CallerFieldName string

	// ErrorFieldName defines the name of the error field.
	ErrorFieldName string

	// ErrorStackFieldName defines the name of the error stack field.
	ErrorStackFieldName string

	// LevelFieldName defines the name of the level field.
	LevelFieldName string

	// MessageFieldName defines the name of the message field.
	MessageFieldName string

	// TimestampFieldName defines the name of the timestamp field.
	TimestampFieldName string
}

// NewJSONWriter creates and initializes a new JSONWriter.
func NewJSONWriter() JSONWriter {
	w := JSONWriter{
		Out: os.Stdout,
	}
	return w
}

// Write handles writing the output to the writer.
func (w JSONWriter) Write(p []byte) (int, error) {
	// decode the JSON message to a generic interface
	var evt map[string]interface{}
	p = decodeIfBinaryToBytes(p)
	d := json.NewDecoder(bytes.NewReader(p))
	d.UseNumber()
	err := d.Decode(&evt)
	if err != nil {
		return 0, fmt.Errorf("cannot decode event: %s", err)
	}

	// update field names
	if _, ok := evt[CallerFieldName]; ok && w.CallerFieldName != "" {
		evt[w.CallerFieldName] = evt[CallerFieldName]
		delete(evt, CallerFieldName)
	}
	if _, ok := evt[ErrorFieldName]; ok && w.ErrorFieldName != "" {
		evt[w.ErrorFieldName] = evt[ErrorFieldName]
		delete(evt, ErrorFieldName)
	}
	if _, ok := evt[ErrorStackFieldName]; ok && w.ErrorStackFieldName != "" {
		evt[w.ErrorStackFieldName] = evt[ErrorStackFieldName]
		delete(evt, ErrorStackFieldName)
	}
	if _, ok := evt[LevelFieldName]; ok && w.LevelFieldName != "" {
		evt[w.LevelFieldName] = evt[LevelFieldName]
		delete(evt, LevelFieldName)
	}
	if _, ok := evt[MessageFieldName]; ok && w.MessageFieldName != "" {
		evt[w.MessageFieldName] = evt[MessageFieldName]
		delete(evt, MessageFieldName)
	}
	if v, ok := evt[TimestampFieldName]; ok {
		if w.TimeFormat != "" {
			if fValue, ok := v.(string); ok {
				if t, err := time.Parse(time.RFC3339, fValue); err == nil {
					evt[TimestampFieldName] = t.Format(w.TimeFormat)
				}
			}
		}
		if w.TimestampFieldName != "" {
			evt[w.TimestampFieldName] = evt[TimestampFieldName]
			delete(evt, TimestampFieldName)
		}
	}

	// remove excluded fields
	for _, field := range w.ExcludeFields {
		delete(evt, field)
	}

	// marshal the output
	b, err := json.Marshal(evt)
	if err != nil {
		return 0, err
	}
	w.Out.Write(b)
	w.Out.Write([]byte("\n"))
	return len(p), err
}
