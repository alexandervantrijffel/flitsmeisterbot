package main

import (
	"runtime/debug"

	"github.com/alexandervantrijffel/flitsmeisterbot/slacklogger"

	"github.com/alexandervantrijffel/flitsmeisterbot/logger"

	"github.com/alexandervantrijffel/go-utilities/filecopier"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.Get().Fatalf("FATAL: %+v (app panicked) %s\n", r, debug.Stack())
		}
	}()
	headers := map[string]string{
		"accept-encoding":           "gzip, deflate, br",
		"accept-language":           "en,en-US;q=0.8,nl;q=0.6",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"cache-control":             "max-age=0",
		"authority":                 "www.flitsmeister.nl",
		"if-modified-since":         "Wed, 01 Nov 2017 09:43:06 GMT",
	}
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
