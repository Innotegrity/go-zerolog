// Package log provides a global logger for zerolog.
package log

import (
	"context"
	"fmt"
	"io"

	"go.innotegrity.dev/zerolog"
)

/** BEGIN CUSTOM CODE */

/*
// Logger is the global logger.
var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
*/

/** END CUSTOM CODE */

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return Logger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return Logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return Logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return Logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return Logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return Logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return Logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return Logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return Logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return Logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return Logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return Logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return Logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return Logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

/** BEGIN CUSTOM CODE */

// ReplaceGlobal replaces the global logger and returns a function to retrieve the original logger.
func ReplaceGlobal(logger zerolog.Logger) (func(), zerolog.Logger) {
	mutex.Lock()
	prev := Logger
	Logger = logger
	mutex.Unlock()
	return func() { ReplaceGlobal(prev) }, logger
}

// GetLevel returns the current log level
func GetLevel() zerolog.Level {
	return Logger.GetLevel()
}

// IsDebugEnabled returns whether or not debug or trace logging is enabled.
func IsDebugEnabled() bool {
	return Logger.GetLevel() <= zerolog.DebugLevel
}

// SetLevel updates the minimum logging level for the logger.
func SetLevel(level zerolog.Level) {
	Logger.SetLevel(level)
}

// WouldLog returns whether or not a mesage with the given level would be logged.
func WouldLog(level zerolog.Level) bool {
	return Logger.WouldLog(level)
}

// InitLogrusInterceptor initializes the logrus global logger with an expected format so that they can be
// intercepted and parsed by ParseLogrusMessages.
func InterceptLogrusMessages(writer *io.PipeWriter) {
	Logger.InterceptLogrusMessages(writer)
}

// ParseLogrusMessages uses a reader paired with the output writer from InitLogrusInterceptor to read messages
// from the pipe in a blocking manner until the writer is closed.
func ParseLogrusMessages(reader *io.PipeReader, handler zerolog.ParseLogrusMessageErrorHandler) {
	Logger.ParseLogrusMessages(reader, handler)
}

// WithContext returns a copy of ctx with l associated. If an instance of Logger
// is already in the context, the context is not updated.
func WithContext(ctx context.Context) context.Context {
	return Logger.WithContext(ctx)
}

/** END CUSTOM CODE */
