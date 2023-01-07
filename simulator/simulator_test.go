package simulator

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/mariajdab/alien-invasion/alien"
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

func TestMainLoop(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	rand.Seed(randomNumberGeneratorSeed)

	aliens := make([]*alien.Alien, 4)

	// fill the alien data needed for the Simulator struct
	for i := 0; i < 4; i++ {
		aliens[i] = &alien.Alien{
			StepsCount: 0,
		}
		aliens[i].InitRandomPosition(mockCitiesLedger)
		fmt.Println(aliens[i])
	}

	s := Simulator{
		Aliens:     aliens,
		CitiesMap:  mockCitiesMap,
		CityLedger: mockCitiesLedger,
	}

	s.MainLoop()
	t.Log(buf.String())

	if s.Aliens[0] != nil && s.Aliens[2] != nil {
		t.Errorf(`Expected Alien 0 and Alien 2 be nil but got alien 0: %v alien 1: %v"`, s.Aliens[0], s.Aliens[2])
	}

	if s.Aliens[1].StepsCount != maxStepCount && s.Aliens[3].StepsCount != maxStepCount {
		t.Errorf(`Expected Alien 1 and Alien 3 reached 10k steps got alien 1: %v alien 3: %v"`, s.Aliens[1].StepsCount, s.Aliens[3].StepsCount)
	}

}
