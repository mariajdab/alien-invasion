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

func main() {
	rand.Seed(time.Now().Unix())
	citiesMap, citiesLedger := file.OpenAndReadCitiesFile("cities_list.txt")
	alienNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

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
