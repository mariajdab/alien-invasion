package main

import "math/rand"

type Alien struct {
	StepsCount int64
	Position   string
}

func (a *Alien) InitRandomPosition(citiesLedger map[int]string) {
	max := len(citiesLedger)
	a.Position = citiesLedger[rand.Intn(max)]
}

func (a *Alien) UpdatePosition(citiesLedger map[int]string, citiesMap map[string]map[string]bool) {
	citiesIds := keysFromMap(citiesLedger)
	newPos := rand.Intn(len(citiesIds))

	for citiesMap[a.Position][citiesLedger[newPos]] {
		newPos = rand.Intn(len(citiesLedger))
	}
	a.Position = citiesLedger[newPos]
}

func keysFromMap(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func remove(slice []*Alien, s int) {
	slice[s] = nil
}
