package alien

import (
	"math/rand"
	"testing"
)

const randomNumberGeneratorSeed = 1971603567

var citiesLedger = map[int]string{0: "Bar", 1: "Baz", 2: "Foo"}

var citiesMap = map[string]map[string]bool{
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

func TestInitRandomPosition(t *testing.T) {
	rand.Seed(randomNumberGeneratorSeed)
	alien := Alien{}
	alien.InitRandomPosition(citiesLedger)
	if alien.Position != "Baz" {
		t.Errorf(`Expected "Baz" but got "%s"`, alien.Position)
	}

	alien.InitRandomPosition(citiesLedger)
	if alien.Position != "Foo" {
		t.Errorf(`Expected "Baz" but got "%s"`, alien.Position)
	}
}

func TestUpdatePosition(t *testing.T) {
	rand.Seed(randomNumberGeneratorSeed)
	alien := Alien{}
	alien.InitRandomPosition(citiesLedger)

	alien.UpdatePosition(citiesLedger, citiesMap)
	if alien.Position != "Baz" {
		t.Errorf(`Expected "Baz" but got "%s"`, alien.Position)
	}

	alien.UpdatePosition(citiesLedger, citiesMap)
	if alien.Position != "Bar" {
		t.Errorf(`Expected "Foo" but got "%s"`, alien.Position)
	}
}
