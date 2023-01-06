package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadCitiesFile(fileName string) (map[string]map[string]bool, map[int]string) {
	var cityNumber int

	citiesMap := make(map[string]map[string]bool)
	citiesLedger := make(map[int]string)

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		mainCity, adjacentCities := processTokens(tokens)

		citiesMap[mainCity] = adjacentCities
		citiesLedger[cityNumber] = mainCity

		cityNumber++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(citiesMap, citiesLedger)

	return citiesMap, citiesLedger
}

func processTokens(tokens []string) (string, map[string]bool) {
	adjacentCities := make(map[string]bool)

	mainCity := tokens[0]

	for _, token := range tokens[1:] {
		adjCities := strings.Split(token, "=")
		adjacentCities[adjCities[1]] = true
	}

	return mainCity, adjacentCities
}
