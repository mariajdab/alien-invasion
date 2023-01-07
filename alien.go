package main

import "math/rand"

type Alien struct {
	StepsCount int64
	Position   string
}

// InitRandomPosition assigns a random position for the alien, the position will be between the numbers
// that are associated with a city (citiesLedger)
func (a *Alien) InitRandomPosition(citiesLedger map[int]string) {
	max := len(citiesLedger)
	a.Position = citiesLedger[rand.Intn(max)]
}

// UpdatePosition
func (a *Alien) UpdatePosition(citiesLedger map[int]string, citiesMap map[string]map[string]bool) {
	citiesIds := keysFromMap(citiesLedger)
	newPos := rand.Intn(len(citiesIds))

	for citiesMap[a.Position][citiesLedger[newPos]] {
		newPos = rand.Intn(len(citiesLedger))
	}
	a.Position = citiesLedger[newPos]
}

// keysFromMap
func keysFromMap(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// remove deletes an alien assigning it to nil
func remove(slice []*Alien, s int) {
	slice[s] = nil
}
