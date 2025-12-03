package main

import (
	"os"
	"github.com/bugsfounder/logger"
)

func main() {
	// Use the default singleton logger
	log := logger.Logging()
	log.Info("This is an info message from the default logger")
	log.Warning("This is a warning message")
	log.Error("This is an error message")

	// Create a custom logger instance
	customLogger := logger.AppLogger(os.Stdout)
	customLogger.Debug("This is a debug message from a custom logger")
}
