package file

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// OpenAndReadCitiesFile opens and reads the cities file and fills citiesMap and citiesLedger maps
func OpenAndReadCitiesFile(fileName string) (map[string]map[string]bool, map[int]string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	return ReadCitiesFile(f)
}

func ReadCitiesFile(reader io.Reader) (map[string]map[string]bool, map[int]string) {
	cityNumber := 0
	// stores each city with his adjacent cities
	citiesMap := make(map[string]map[string]bool)
	// associates a number with a city
	citiesLedger := make(map[int]string)
	scanner := bufio.NewScanner(reader)

	// reads each line, split it and fill the maps
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

	return citiesMap, citiesLedger
}

// processTokens retrieves each useful information (mainCity and adjCities) of the line that was split with spaces
func processTokens(tokens []string) (string, map[string]bool) {
	adjacentCities := make(map[string]bool)

	mainCity := tokens[0]

	// ignores the first item of the list (mainCity)
	for _, token := range tokens[1:] {
		adjCities := strings.Split(token, "=")
		// ignores the direction
		adjacentCities[adjCities[1]] = true
	}

	return mainCity, adjacentCities
}
