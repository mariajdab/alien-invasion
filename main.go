package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mariajdab/alien-invasion/alien"
	"github.com/mariajdab/alien-invasion/file"
)

const maxStepCount = 10000
const aliensFight = 2

type Simulator struct {
	aliens     []*alien.Alien
	citiesMap  map[string]map[string]bool
	cityLedger map[int]string
}

func main() {
	rand.Seed(time.Now().Unix())
	citiesMap, citiesLedger := file.ReadCitiesFile("cities_list.txt")
	//alienNumber, err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//
	//}

	alienNumber := 2
	aliens := make([]*alien.Alien, 2)

	for i := 0; i < alienNumber; i++ {
		aliens[i] = &alien.Alien{
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

// mainLoop starts and maintains the simulator until the conditional is not valid
func (s Simulator) mainLoop() {
	// each alien hasn't moved at least 10k times or all the aliens haven't been destroyed,
	// also it's considered that the citiesMap could be a nil map, so the condition is added to prevent an error
	for aliensUnder10kOrAllNil(s.aliens) && s.citiesMap != nil {
		log.Println(s.aliens, s.cityLedger, s.citiesMap)

		s.updateMap()
		s.updateAliensInMap()

		time.Sleep(1 * time.Second)
	}
}

// updateMap updates the map removing the city from the map if two aliens are in the city, also the two aliens will be destroyed
func (s Simulator) updateMap() {
	aliensPerCity := make(map[string][]int)

	for index, alien := range s.aliens {
		if alien != nil {
			log.Printf("alient: %v index: %d", alien, index)
			aliensPerCity[alien.Position] = append(aliensPerCity[alien.Position], index)
		}
	}

	for city, aliensIndexes := range aliensPerCity {
		fmt.Println(aliensIndexes)
		if len(aliensIndexes) == aliensFight {
			log.Printf("%s has been destryed by alien %d and alien %d\n", city, aliensIndexes[0], aliensIndexes[1])
			delete(s.citiesMap, city)

			alien.Remove(s.aliens, aliensIndexes[0])
			alien.Remove(s.aliens, aliensIndexes[1])

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

// aliensUnder10kOrAllNi in charges of update the alien position for each iteration
func (s Simulator) updateAliensInMap() {
	for _, alien := range s.aliens {
		if alien != nil {
			alien.StepsCount++
			alien.UpdatePosition(s.cityLedger, s.citiesMap)
		}
	}
}

// aliensUnder10kOrAllNi will return false if all the aliens are nil or all the aliens has reached 10k steps
func aliensUnder10kOrAllNil(aliens []*alien.Alien) bool {
	for _, alien := range aliens {
		if alien != nil && alien.StepsCount < maxStepCount {
			return true
		}
	}

	return false
}
