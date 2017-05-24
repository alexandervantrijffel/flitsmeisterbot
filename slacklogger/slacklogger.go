package slacklogger

import (
	"os"

	"github.com/alexandervantrijffel/flitsmeisterbot/flags"

	log "github.com/Sirupsen/logrus"
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"
)

func Get(isError bool) *log.Logger {
	logger := log.New()
	logger.Out = os.Stdout
	logger.Hooks.Add(GetHook(isError, log.DebugLevel))
	return logger
}

func GetHook(isError bool, minLevel log.Level) log.Hook {
	emoji := ":racing_car:"
	if isError {
		emoji = ":fire:"
	}
	cfg := lrhook.Config{
		MinLevel: minLevel,
		Message: chat.Message{
			IconEmoji: emoji,
		},
	}
	hook := lrhook.New(cfg, flags.Flags.SlackWebhookURL)
	return hook
}
