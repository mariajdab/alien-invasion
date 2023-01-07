package simulator

import (
	"log"

	"github.com/mariajdab/alien-invasion/alien"
)

const maxStepCount = 10000
const aliensFight = 2

type Simulator struct {
	Aliens     []*alien.Alien
	CitiesMap  map[string]map[string]bool
	CityLedger map[int]string
}

// MainLoop starts and maintains the simulator until the conditional is not valid
func (s *Simulator) MainLoop() {
	// each alien hasn't moved at least 10k times or all the Aliens haven't been destroyed,
	// also it's considered that the CitiesMap could be a nil map, so the condition is added to prevent an error
	for s.aliensUnder10kOrAllNil() && s.CitiesMap != nil {
		s.updateMap()
		s.updateAliensInMap()
	}
}

// updateMap updates the map removing the city from the map if two Aliens are in the city, also the two Aliens will be destroyed
func (s *Simulator) updateMap() {
	aliensPerCity := make(map[string][]int)

	for index, alien := range s.Aliens {
		if alien != nil {
			aliensPerCity[alien.Position] = append(aliensPerCity[alien.Position], index)
			//log.Printf("%#v\n", alien)
		}
	}

	for city, aliensIndexes := range aliensPerCity {
		if len(aliensIndexes) == aliensFight {
			log.Printf("%s has been destryed by alien %d and alien %d\n", city, aliensIndexes[0], aliensIndexes[1])
			delete(s.CitiesMap, city)

			alien.Remove(s.Aliens, aliensIndexes[0])
			alien.Remove(s.Aliens, aliensIndexes[1])

			for _, adjacentCities := range s.CitiesMap {
				delete(adjacentCities, city)
			}

			for cityID, cityName := range s.CityLedger {
				if cityName == city {
					delete(s.CityLedger, cityID)
					break
				}
			}
		}
	}
}

// aliensUnder10kOrAllNi in charges of update the alien position for each iteration
func (s *Simulator) updateAliensInMap() {
	for _, alien := range s.Aliens {
		if alien != nil {
			alien.StepsCount++
			alien.UpdatePosition(s.CitiesMap)
		}
	}
}

// aliensUnder10kOrAllNi will return false if all the Aliens are nil or all the Aliens has reached 10k steps
func (s *Simulator) aliensUnder10kOrAllNil() bool {
	for _, alien := range s.Aliens {
		if alien != nil && alien.StepsCount < maxStepCount {
			return true
		}
	}

	return false
}
