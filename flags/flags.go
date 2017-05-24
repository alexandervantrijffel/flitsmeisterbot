package flags

import (
	"flag"
	"fmt"
	"os"
)

var Flags AppFlags = installFlags()

type AppFlags struct {
	SlackWebhookURL string
	LogLevel        string
}

func installFlags() AppFlags {
	var a AppFlags
	flag.StringVar(&a.SlackWebhookURL, "slackurl", "", "Set the URL for posting to Slack using an incoming webhook. Example: https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")
	flag.StringVar(&a.LogLevel, "log-level", "debug", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	flag.Parse()
	if a.SlackWebhookURL == "" {
		fmt.Println("The slackurl for posting to a Slack inoming webhook is required.")
		os.Exit(1)
	}
	return a
}
