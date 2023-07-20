package log

/** BEGIN CUSTOM CODE */

import (
	"os"
	"sync"

	"go.innotegrity.dev/zerolog"
)

var (
	mutex sync.RWMutex

	// Logger is the global logger.
	Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
)

/** END CUSTOM CODE */
