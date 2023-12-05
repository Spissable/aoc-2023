package day5

import (
	"aoc-2023/util"
	"math"
	"strings"
)

type Mapping struct {
	source      int
	destination int
	r           int
}

type Almanac struct {
	seeds                 []int
	seedToSoil            []Mapping
	soilToFertilizer      []Mapping
	fertilizerToWater     []Mapping
	waterToLight          []Mapping
	lightToTemperature    []Mapping
	temperatureToHumidity []Mapping
	humidityToLocation    []Mapping
}

func solvePuzzle1(almanac Almanac) int {
	lowestLocation := math.MaxInt

	for _, seed := range almanac.seeds {
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

	return lowestLocation
}

func getProperty(mappings []Mapping, source int) int {
	for _, m := range mappings {
		if source >= m.source && source < m.source+m.r {
			return m.destination + source - m.source
		}
	}

	return source
}

func setProperty(almanac Almanac, header string, source, destination, r int) Almanac {
	m := Mapping{
		source:      source,
		destination: destination,
		r:           r,
	}
	if strings.HasPrefix(header, "seed-to-soil") {
		almanac.seedToSoil = append(almanac.seedToSoil, m)
	} else if strings.HasPrefix(header, "soil-to-fertilizer") {
		almanac.soilToFertilizer = append(almanac.soilToFertilizer, m)
	} else if strings.HasPrefix(header, "fertilizer-to-water") {
		almanac.fertilizerToWater = append(almanac.fertilizerToWater, m)
	} else if strings.HasPrefix(header, "water-to-light") {
		almanac.waterToLight = append(almanac.waterToLight, m)
	} else if strings.HasPrefix(header, "light-to-temperature") {
		almanac.lightToTemperature = append(almanac.lightToTemperature, m)
	} else if strings.HasPrefix(header, "temperature-to-humidity") {
		almanac.temperatureToHumidity = append(almanac.temperatureToHumidity, m)
	} else if strings.HasPrefix(header, "humidity-to-location") {
		almanac.humidityToLocation = append(almanac.humidityToLocation, m)
	}
	return almanac
}

func parseAlmanac(input string) Almanac {
	almanac := Almanac{}

	lines := strings.Split(input, "\n")
	i := 0

	for {
		if strings.HasPrefix(lines[i], "seeds:") {
			almanac.seeds = util.StringsToNums(strings.Split(strings.Trim(strings.Split(lines[i], ":")[1], " "), " "))
		} else if len(lines[i]) > 0 {
			header := lines[i]
			i++
			for len(lines[i]) > 0 {
				nums := util.StringsToNums(strings.Split(lines[i], " "))
				destination := nums[0]
				source := nums[1]
				r := nums[2]

				almanac = setProperty(almanac, header, source, destination, r)

				if i == len(lines)-1 {
					return almanac
				}

				i++
			}
		}

		i++
	}
}
