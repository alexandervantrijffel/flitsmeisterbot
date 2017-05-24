package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunEvaluationIntegration(t *testing.T) {
	_, newSpeedtraps := runSpeedtrapsEvaluation("dataunittest/current.geojson", "dataunittest/previous.geojson")
	assert.Equal(t, []speedtrap{
		speedtrap{road: "A12", location: "10.0 (Zoetermeer)"},
	}, newSpeedtraps)
}

func TestLoadSpeedtraps(t *testing.T) {
	features := loadAllSpeedtrapsInNL("dataunittest/previous.geojson")
	assert.True(t, len(features) > 0, "should have parsed features")
	for _, f := range features {
		assert.Equal(t, "speedtrap", f.Properties.TypeDescription, "all features should be of type speedtrap")
		assert.Equal(t, "nl", f.Properties.CountryCode, "All speedtraps should be in NL")
	}
	formattedLines := featuresToSpeedtraps(features)
	assert.Equal(t, []speedtrap{
		speedtrap{road: "A2", location: "180.0 (Leende)"},
		speedtrap{road: "N7", location: "199.2 (Groningen)"},
		speedtrap{road: "N302", location: "86.6 (Lelystad)"},
	}, formattedLines)
}

func alertForAllRoads() []string {
	return []string{"A1", "A2", "A3", "A4", "A5", "A6"}
}

func TestShouldNotReportNotPreferredRoad(t *testing.T) {
	reported := reportNewSpeedtraps(
		alertForAllRoads(),
		[]speedtrap{testtrap("N1")},
		[]speedtrap{})
	assert.Equal(t, 0, len(reported))
}

func TestShouldReportOnlyNewSpeedtraps(t *testing.T) {
	reported := reportNewSpeedtraps(
		alertForAllRoads(),
		[]speedtrap{testtrap("A1"), testtrap("A2"), testtrap("A3")},
		[]speedtrap{testtrap("A5"), testtrap("A1"), testtrap("A3")})
	assert.Equal(t, 1, len(reported))
	assert.Equal(t, "A2", reported[0].road)
}

func TestShouldReportAllSpeedtraps(t *testing.T) {
	reported := reportNewSpeedtraps(
		alertForAllRoads(),
		[]speedtrap{testtrap("A1"), testtrap("A2")},
		[]speedtrap{})
	assert.Equal(t, 2, len(reported))
}

func TestShouldNoSpeedtrapsWhenNoNewSpeedtraps(t *testing.T) {
	reported := reportNewSpeedtraps(
		alertForAllRoads(),
		[]speedtrap{},
		[]speedtrap{testtrap("A1"), testtrap("A2")})
	assert.Equal(t, 0, len(reported))
}

func testtrap(road string) speedtrap {
	return speedtrap{road: road, location: road}
}
