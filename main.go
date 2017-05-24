package main

import (
	"0local/alexandervantrijffel/flitsmeisterbot/logger"
	"0local/alexandervantrijffel/flitsmeisterbot/slacklogger"
	"runtime/debug"

	"github.com/alexandervantrijffel/go-utilities/filecopier"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.Get().Fatalf("FATAL: %+v (app panicked) %s\n", r, debug.Stack())
		}
	}()
	headers := map[string]string{
		"accept":           "application/json, text/javascript, */*; q=0.01",
		"accept-encoding":  "gzip, deflate, sdch, br",
		"accept-language":  "en,en-US;q=0.8,nl;q=0.6",
		"cache-control":    "no-cache",
		"pragma":           "no-cache",
		"referer":          "https://www.flitsmeister.nl/kaart.html",
		"user-agent":       "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"x-requested-with": "XMLHttpRequest"}
	current := "data/current.geojson"
	previous := "data/previous.geojson"
	curlURL("https://www.flitsmeister.nl/assets/data/current.geojson", headers, current)
	currentSpeedtraps, newSpeedtraps := runSpeedtrapsEvaluation(current, previous)
	logger.Get().Infof("found NL speedtraps (reporting %d) %s\n", len(newSpeedtraps), currentSpeedtraps)
	for _, f := range newSpeedtraps {
		slacklogger.Get(false).Infof("!!Speed trap alert!! %s", f)
	}
	check(filecopier.CopyFile(current, previous))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
