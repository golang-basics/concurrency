package main

import (
	"fmt"
	"log"
	"os"
)

func loggerInit() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetPrefix("[LOG] ")
}

func logInfo(format string, args ...interface{}) {
	logWithPrefix("INFO", format, args)
}

func logDebug(format string, args ...interface{}) {
	logWithPrefix("DEBUG", format, args)
}

func logError(format string, args ...interface{}) {
	logWithPrefix("ERROR", format, args)
}

func logWithPrefix(prefix, format string, args []interface{}) {
	log.SetPrefix(fmt.Sprintf("[%s] ", prefix))
	log.Printf(format, args...)
}
