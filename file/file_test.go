package file

import (
	"bytes"
	"testing"
)

const mockFileData = `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Baz east=Foo
Qu-ux north=Foo
Bee east=Bar
`

var (
	mockCitiesLedger = map[int]string{0: "Foo", 1: "Bar", 2: "Baz", 3: "Qu-ux", 4: "Bee"}
	mockCitiesMap    = map[string]map[string]bool{
		"Foo": {
			"Bar":   true,
			"Baz":   true,
			"Qu-ux": true,
		},
		"Bar": {
			"Foo": true,
			"Bee": true,
		},
		"Baz": {
			"Foo": true,
		},
		"Qu-ux": {
			"Foo": true,
		},
		"Bee": {
			"Bar": true,
		},
	}
)

func TestReadCitiesFile(t *testing.T) {
	buffer := bytes.Buffer{}
	buffer.WriteString(mockFileData)
	citiesMap, citiesLedger := ReadCitiesFile(&buffer)
	if !equalMaps(citiesMap, mockCitiesMap) {
		t.Errorf("Invalid file read, expected map %#v, but got %#v", mockCitiesMap, citiesMap)
	}

	if !equalLedgers(citiesLedger, mockCitiesLedger) {
		t.Errorf("Invalid file read, expected ledger %#v, but got %#v", mockCitiesLedger, citiesLedger)
	}
}

func equalMaps(map1, map2 map[string]map[string]bool) bool {
	for city, adjacentCities := range map1 {
		for adjacentCity := range adjacentCities {
			if !map2[city][adjacentCity] {
				return false
			}
		}
	}

	return true
}

func equalLedgers(ledger1, ledger2 map[int]string) bool {
	for cityID, cityName := range ledger1 {
		if ledger2[cityID] != cityName {
			return false
		}
	}

	return true
}
