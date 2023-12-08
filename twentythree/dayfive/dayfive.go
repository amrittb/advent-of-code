package dayfive

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

var shouldLog = false

func parseInitialSeeds(line string) []int {
  splitLine := strings.Split(line, ":")
  if len(splitLine) != 2 {
    log.Fatalln("Invalid Inital Seeds Line: " + line)
  }

  return convertToIntSlice(splitLine[1], " ")
}

func FindLowestLocation(lines []string) int {
  seedsPattern := regexp.MustCompile("seeds:.*")
  aTobPattern := regexp.MustCompile("([a-z]+)-to-([a-z]+) map:")
  
  var initialSeeds []int

  seedToSoil := [][]int{}
  soilToFertilizer := [][]int{}
  fertilizerToWater := [][]int{}
  waterToLight := [][]int{}
  lightToTemp := [][]int{}
  tempToHumidity := [][]int{}
  humidityToLocation := [][]int{}

  var selectedSlice *([][]int)
  for _, line := range lines {
    line = strings.Trim(line, " ")
    if line == "" {
      continue
    }

    if seedsPattern.MatchString(line) {
      initialSeeds = parseInitialSeeds(line)
      continue
    }

    matches := aTobPattern.FindStringSubmatch(line)
    if len(matches) != 0 {
      if len(matches) != 3 {
	log.Fatalf("Matches should be equal to three: %v\n", matches)
      }

      // Check a to b map
      mapPair := matches[1] + "-" + matches[2]
      switch mapPair {
      case "seed-soil":
	selectedSlice = &seedToSoil
      case "soil-fertilizer":
	selectedSlice = &soilToFertilizer
      case "fertilizer-water":
	selectedSlice = &fertilizerToWater
      case "water-light":
	selectedSlice = &waterToLight
      case "light-temperature":
	selectedSlice = &lightToTemp
      case "temperature-humidity":
	selectedSlice = &tempToHumidity
      case "humidity-location":
	selectedSlice = &humidityToLocation
      default:
	log.Fatalf("Map Pair: %v is invalid\n", mapPair)
      }

      continue
    }

    if selectedSlice == nil {
      log.Fatalf("Selected Slice should not be nil, but is.\n")
    }
    // Didn't match, so parse numbers
    nums := convertToIntSlice(line, " ")
    if len(nums) != 3 {
      log.Fatalf("Parsed numbers should be equal to three: %v\n", nums)
    }

    (*selectedSlice) = append((*selectedSlice), nums)
  }

  lowestLocation := -1

  for _, seed := range initialSeeds {
    soil := convertValue(seedToSoil, seed)
    fertilizer := convertValue(soilToFertilizer, soil)
    water := convertValue(fertilizerToWater, fertilizer)
    light := convertValue(waterToLight, water)
    temp := convertValue(lightToTemp, light)
    humidity := convertValue(tempToHumidity, temp)
    location := convertValue(humidityToLocation, humidity)

    if shouldLog {
      fmt.Printf("seed: %v\n", seed)
      fmt.Printf("soil: %v\n", soil)
      fmt.Printf("fertilizer: %v\n", fertilizer)
      fmt.Printf("water: %v\n", water)
      fmt.Printf("light: %v\n", light)
      fmt.Printf("temp: %v\n", temp)
      fmt.Printf("humidity: %v\n", humidity)
      fmt.Printf("location: %v\n", location)
    }

    if lowestLocation == -1 {
      lowestLocation = location
    }
    lowestLocation = integer.Min(lowestLocation, location)
  }

  return lowestLocation
}

func convertToIntSlice(s, sep string) []int {
  ints := []int{}
  intStrs := strings.Split(strings.Trim(s, " "), sep) 
  for _, intStr := range intStrs {
    i, err := strconv.Atoi(intStr)
    if err != nil {
      log.Fatalln(err)
    }

    ints = append(ints, i)
  }

  return ints
}



func convertValue(conversionChart [][]int, val int) int {
  for _, chart := range conversionChart {
    des, src, r := chart[0], chart[1], chart[2]

    if val >= src && val <= (src + r) {
      return des + (val - src)
    }
  }

  return val
}
