package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func SetLogLevel() {
	LogLevel = strings.ToLower(LogLevel)
	switch LogLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
		log.Trace("Trace log level enabled")
		return
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug log level enabled")
		return
	case "warn":
		log.SetLevel(log.WarnLevel)
		return
	case "error":
		log.SetLevel(log.ErrorLevel)
		return
	}
	log.SetLevel(log.InfoLevel)
}
