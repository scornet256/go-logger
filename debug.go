package logger

import (
	"io"
	"log"
)

var debug bool

// Enables debug logger if set to true. Default is false.
func SetDebug(debugSetting bool) {
	debug = debugSetting
	if !debug {
		log.SetOutput(io.Discard)
	}
}

// Returns the current debug value as a boolean.
func GetDebug() bool {
	return debug
}
