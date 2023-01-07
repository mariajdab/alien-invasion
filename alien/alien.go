package alien

import "math/rand"

type Alien struct {
	StepsCount int64
	Position   string
}

// InitRandomPosition assigns a random position for the alien, the position will be between the numbers
// that are associated with a city (mockCitiesLedger)
func (a *Alien) InitRandomPosition(citiesLedger map[int]string) {
	max := len(citiesLedger)
	a.Position = citiesLedger[rand.Intn(max)]
}

// UpdatePosition
func (a *Alien) UpdatePosition(citiesMap map[string]map[string]bool) {
	citiesIds := keysFromMap(citiesMap[a.Position])
	if len(citiesIds) > 0 {
		a.Position = citiesIds[rand.Intn(len(citiesIds))]
	}
}

// keysFromMap
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
