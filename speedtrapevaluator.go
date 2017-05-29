package main

import (
	"fmt"
	"os"

	"github.com/Mparaiso/lodash-go"
)

type speedtrap struct {
	road     string
	location string
}

func (s speedtrap) String() string {
	return fmt.Sprintf("%s %s", s.road, s.location)
}

func runSpeedtrapsEvaluation(currentFilePath string, previousFilePath string) (currentSpeedtraps []string, newSpeedtraps []speedtrap) {
	currentFeatures := featuresToSpeedtraps(loadAllSpeedtrapsInNL(currentFilePath))
	var previousFeatures []speedtrap
	if _, err := os.Stat(previousFilePath); err == nil || !os.IsNotExist(err) {
		previousFeatures = featuresToSpeedtraps(loadAllSpeedtrapsInNL(previousFilePath))
	}
	check(lo.Map(currentFeatures, func(f speedtrap) string { return f.road + " " + f.location }, &currentSpeedtraps))
	alertForRoads := []string{"A12", "A13", "A16", "A20", "N219", "N210"}
	newSpeedtraps = reportNewSpeedtraps(alertForRoads, currentFeatures, previousFeatures)
	return currentSpeedtraps, newSpeedtraps
}

func featuresToSpeedtraps(speedtraps []flitsmeisterFeature) []speedtrap {
	var theSpeedTraps []speedtrap
	for _, f := range speedtraps {
		theSpeedTraps = append(theSpeedTraps, speedtrap{road: f.Properties.Road,
			location: fmt.Sprintf("%.1f (%s)", f.Properties.Hmp, f.Properties.Location)})
	}
	return theSpeedTraps
}

func reportNewSpeedtraps(alertForRoads []string, speedtraps []speedtrap, previousSpeedtraps []speedtrap) []speedtrap {
	var newSpeedtraps []speedtrap
	err := lo.
		In(speedtraps).
		Filter(func(speedtrap speedtrap) bool {
			isPreferred, _ := lo.IndexOf(alertForRoads, speedtrap.road, 0)
			wasInPrevious, _ := lo.IndexOf(previousSpeedtraps, speedtrap.road, 0)
			return isPreferred != -1 && wasInPrevious == -1
		}).
		Out(&newSpeedtraps)
	check(err)
	return newSpeedtraps
}
