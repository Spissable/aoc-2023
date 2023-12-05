package day5

import (
	"math"
)

func solvePuzzle2(almanac Almanac) int {
	lowestLocation := math.MaxInt
	var channels []chan int

	for i := 0; i < len(almanac.seeds); i += 2 {
		start := almanac.seeds[i]
		r := almanac.seeds[i+1]

		result := make(chan int)
		channels = append(channels, result)

		go getRangeMin(almanac, start, r, result)
	}

	for _, channel := range channels {
		result := <-channel

		if result < lowestLocation {
			lowestLocation = result
		}
	}

	return lowestLocation
}

func getRangeMin(almanac Almanac, start, r int, result chan int) {
	lowestLocation := math.MaxInt

	for seed := start; seed < start+r; seed++ {
		soil := getProperty(almanac.seedToSoil, seed)
		fertilizer := getProperty(almanac.soilToFertilizer, soil)
		water := getProperty(almanac.fertilizerToWater, fertilizer)
		light := getProperty(almanac.waterToLight, water)
		temperature := getProperty(almanac.lightToTemperature, light)
		humidity := getProperty(almanac.temperatureToHumidity, temperature)
		location := getProperty(almanac.humidityToLocation, humidity)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	result <- lowestLocation
}
