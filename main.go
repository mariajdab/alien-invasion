package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mariajdab/alien-invasion/alien"
	"github.com/mariajdab/alien-invasion/file"
	"github.com/mariajdab/alien-invasion/simulator"
)

const worldMap = "cities_list.txt"

func main() {
	rand.Seed(time.Now().UnixNano())

	// the file name is an assumption
	citiesMap, citiesLedger, err := file.OpenAndReadCitiesFile(worldMap)
	if err != nil {
		log.Fatalln(err)
	}

	// number of aliens through a command-line argument
	alienNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// fill the data needed for each alien with stepCount 0
	aliens := make([]*alien.Alien, alienNumber)
	for i := 0; i < alienNumber; i++ {
		aliens[i] = &alien.Alien{
			StepsCount: 0,
		}
		aliens[i].InitRandomPosition(citiesLedger)
	}

	s := simulator.Simulator{
		Aliens:     aliens,
		CitiesMap:  citiesMap,
		CityLedger: citiesLedger,
	}

	s.MainLoop()
}
