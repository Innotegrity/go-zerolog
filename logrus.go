package zerolog

/** BEGIN CUSTOM CODE */
import (
	"github.com/sirupsen/logrus"
)

// logrusMessage represents a JSON-formatted log message.
type logrusMessage struct {
	Level   string                 `json:"level"`
	Message string                 `json:"msg"`
	Method  string                 `json:"func"`
	Data    map[string]interface{} `json:"data"`
}

// loggerLevelToLoggerLevel converts our logger logging levels to logrus logging levels.
func loggerLevelToLogrusLevel(l Level) logrus.Level {
	switch l {
	case PanicLevel:
		return logrus.PanicLevel
	case FatalLevel:
		return logrus.FatalLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case WarnLevel:
		return logrus.WarnLevel
	case InfoLevel:
		return logrus.InfoLevel
	case DebugLevel:
		return logrus.DebugLevel
	case TraceLevel:
		return logrus.TraceLevel
	}
	return 0
}

// logrusLevelToLoggerLevel converts logrus logging levels to our logger logging levels.
func logrusLevelToLoggerLevel(l logrus.Level) Level {
	switch l {
	case logrus.PanicLevel:
		return PanicLevel
	case logrus.FatalLevel:
		return FatalLevel
	case logrus.ErrorLevel:
		return ErrorLevel
	case logrus.WarnLevel:
		return WarnLevel
	case logrus.InfoLevel:
		return InfoLevel
	case logrus.DebugLevel:
		return DebugLevel
	case logrus.TraceLevel:
		return TraceLevel
	}
	return NoLevel
}

/** END CUSTOM CODE */
