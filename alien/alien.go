package alien

import "math/rand"

type Alien struct {
	StepsCount int64
	Position   string
}

// InitRandomPosition assigns the initial random position for the alien, the position will be between the numbers
// that are associated with a city (mockCitiesLedger)
func (a *Alien) InitRandomPosition(citiesLedger map[int]string) {
	max := len(citiesLedger)
	a.Position = citiesLedger[rand.Intn(max)]
}

// UpdatePosition update the alien position when generating a random int where the max number is adjacent cities len
func (a *Alien) UpdatePosition(citiesMap map[string]map[string]bool) {
	citiesIDs := keysFromMap(citiesMap[a.Position])
	// check if the main city has any adjacent city
	if len(citiesIDs) > 0 {
		a.Position = citiesIDs[rand.Intn(len(citiesIDs))]
	}
}

// keysFromMap retrieves the slice of adjacent cities
func keysFromMap(m map[string]bool) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Remove remove deletes an alien assigning it to nil
func Remove(slice []*Alien, s int) {
	slice[s] = nil
}
