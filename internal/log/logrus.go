package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)
}

// Logger Return the singleton log
func Logger() *log.Logger {
	return logger
}
