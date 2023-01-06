package main

import (
	"log"
	"math/rand"
	"time"
)

const maxStepCount = 10000

type Simulator struct {
	aliens     []*Alien
	citiesMap  map[string]map[string]bool
	cityLedger map[int]string
}

func main() {
	rand.Seed(time.Now().Unix())
	citiesMap, citiesLedger := ReadCitiesFile("cities_list.txt")
	//alienNumber, err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//
	//}

	alienNumber := 2
	aliens := make([]*Alien, 2)

	for i := 0; i < alienNumber; i++ {
		aliens[i] = &Alien{
			StepsCount: 0,
		}
		aliens[i].InitRandomPosition(citiesLedger)
	}

	s := Simulator{
		aliens:     aliens,
		citiesMap:  citiesMap,
		cityLedger: citiesLedger,
	}

	s.mainLoop()
}

func (s Simulator) mainLoop() {

	for aliensUnder10kOrAllNil(s.aliens) && s.citiesMap != nil {
		log.Println(s.aliens, s.cityLedger, s.citiesMap, "este")

		s.updateMap()
		s.updateAliensInMap()

		time.Sleep(1 * time.Second)
	}
}

func (s Simulator) updateMap() {
	aliensPerCity := make(map[string][]int)

	for index, alien := range s.aliens {
		if alien != nil {
			aliensPerCity[alien.Position] = append(aliensPerCity[alien.Position], index)
		}
	}

	for city, aliensIndexes := range aliensPerCity {
		if len(aliensIndexes) == 2 {
			log.Printf("%s has been destryed by alien %d and alien %d\n", city, aliensIndexes[0], aliensIndexes[1])
			delete(s.citiesMap, city)

			remove(s.aliens, aliensIndexes[0])
			remove(s.aliens, aliensIndexes[1])

			for _, adjacentCities := range s.citiesMap {
				delete(adjacentCities, city)
			}

			cityIDToDelete := 0
			for cityID, cityName := range s.cityLedger {
				if cityName == city {
					cityIDToDelete = cityID
				}
			}

			delete(s.cityLedger, cityIDToDelete)
		}
	}
}

func (s Simulator) updateAliensInMap() {
	for _, alien := range s.aliens {
		if alien != nil {
			alien.StepsCount++
			alien.UpdatePosition(s.cityLedger, s.citiesMap)
		}
	}
}

func aliensUnder10kOrAllNil(aliens []*Alien) bool {
	for _, alien := range aliens {
		if alien != nil && alien.StepsCount < maxStepCount {
			return true
		}
	}

	return false
}
