package logger

import (
	"0local/alexandervantrijffel/flitsmeisterbot/slacklogger"
	"fmt"
	"os"

	"github.com//flitsmeisterbot/flags"

	log "github.com/Sirupsen/logrus"
)

var defaultLogger *log.Entry = newLogger()

func Get() *log.Entry {
	return defaultLogger
}

func newLogger() *log.Entry {
	log.SetOutput(os.Stdout)
	log.AddHook(slacklogger.GetHook(true, log.WarnLevel))
	SetLogLevel(flags.Flags.LogLevel)
	return log.WithFields(log.Fields{
		"app": "flitsmeisterbot",
	})
}

func SetLogLevel(logLevel string) {
	if logLevel != "" {
		lvl, err := log.ParseLevel(logLevel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse logging level: %s\n", logLevel)
			os.Exit(1)
		}
		log.SetLevel(lvl)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
