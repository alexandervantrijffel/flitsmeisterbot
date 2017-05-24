package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// generated with https://mholt.github.io/json-to-go/
// blog: https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
type flitsmeisterFeature struct {
	Type     string `json:"type"`
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"geometry"`
	Properties struct {
		ID              string  `json:"id"`
		CountryCode     string  `json:"country_code"`
		Road            string  `json:"road"`
		RoadLetter      string  `json:"road_letter"`
		Hmp             float64 `json:"hmp,omitempty"`
		TypeID          int     `json:"type_id,omitempty"`
		TypeDescription string  `json:"type_description"`
		Location        string  `json:"location"`
	} `json:"properties"`
}

type flitsmeisterFeatures struct {
	Type     string                `json:"type"`
	Features []flitsmeisterFeature `json:"features"`
}

func loadAllSpeedtrapsInNL(filename string) []flitsmeisterFeature {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	fromJSON, err := ioutil.ReadAll(f)
	check(err)
	var features flitsmeisterFeatures
	if err := json.Unmarshal(fromJSON, &features); err != nil {
		details := fmt.Sprintf("Failed to unmarshal to type flitsmeisterFeatures from content in file %s", filename)
		panic(errors.Wrap(err, details))
	}
	return features.getSpeedTrapsInNL()
}

func (features *flitsmeisterFeatures) getSpeedTrapsInNL() []flitsmeisterFeature {
	var filteredFeatures []flitsmeisterFeature
	for _, f := range features.Features {
		if f.Properties.TypeDescription == "speedtrap" && f.Properties.CountryCode == "nl" {
			filteredFeatures = append(filteredFeatures, f)
		}
	}
	return filteredFeatures
}
