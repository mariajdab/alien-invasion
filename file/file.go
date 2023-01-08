package file

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// OpenAndReadCitiesFile opens and reads the cities file and fills citiesMap and citiesLedger maps
func OpenAndReadCitiesFile(fileName string) (map[string]map[string]bool, map[int]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()
	return readCitiesFile(f)
}

func readCitiesFile(reader io.Reader) (map[string]map[string]bool, map[int]string, error) {
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
		return nil, nil, err
	}
	return citiesMap, citiesLedger, nil
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
