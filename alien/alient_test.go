package alien

import (
	"math/rand"
	"testing"
)

const randomNumberGeneratorSeed = 1971603567

var (
	mockCitiesLedger = map[int]string{0: "Bar", 1: "Baz", 2: "Foo"}

	mockCitiesMap = map[string]map[string]bool{
		"Foo": {
			"Bar":   true,
			"Baz":   true,
			"Qu-ux": true,
		},
		"Bar": {
			"Foo": true,
			"Bee": true,
		},
		"Bee": {
			"Bar": true,
		},
		"Baz": {
			"Foo": true,
		},
		"Qu-ux": {
			"Foo": true,
		},
	}
)

func TestInitRandomPosition(t *testing.T) {
	// the seed will allow to handle the test due it prevents an unexpected behaviour
	rand.Seed(randomNumberGeneratorSeed)

	alien := Alien{}
	alien.InitRandomPosition(mockCitiesLedger)
	if alien.Position != "Baz" {
		t.Errorf(`Expected "Baz" but got "%s"`, alien.Position)
	}
}

func TestUpdatePosition(t *testing.T) {
	// the seed will allow to handle the test due it prevents an unexpected behaviour
	rand.Seed(randomNumberGeneratorSeed)

	alien := Alien{}
	alien.InitRandomPosition(mockCitiesLedger)

	alien.UpdatePosition(mockCitiesMap)
	if alien.Position != "Foo" {
		t.Errorf(`Expected "Foo" but got "%s"`, alien.Position)
	}
}
