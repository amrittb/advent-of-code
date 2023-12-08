package dayfive

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

// BruteForce
// TODO: Fix Later
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
    nums := integer.ConvertToIntSlice(line, " ")
    if len(nums) != 3 {
      log.Fatalf("Parsed numbers should be equal to three: %v\n", nums)
    }

    (*selectedSlice) = append((*selectedSlice), nums)
  }

  lowestLocation := -1

  for i := 0; i < len(initialSeeds); i++ {
    soil := convertValue(seedToSoil, initialSeeds[i])
    fertilizer := convertValue(soilToFertilizer, soil)
    water := convertValue(fertilizerToWater, fertilizer)
    light := convertValue(waterToLight, water)
    temp := convertValue(lightToTemp, light)
    humidity := convertValue(tempToHumidity, temp)
    location := convertValue(humidityToLocation, humidity)

    if lowestLocation == -1 {
      lowestLocation = location
    }
    lowestLocation = integer.Min(lowestLocation, location)
  }

  return lowestLocation
}

// BruteForce
// TODO: Fix Later
func FindLowestLocationOfSeedRanges(lines []string) int {
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
      if len(initialSeeds) % 2 != 0 {
	log.Fatalf("Initial seed count is not even: %v", len(initialSeeds))
      }
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
    nums := integer.ConvertToIntSlice(line, " ")
    if len(nums) != 3 {
      log.Fatalf("Parsed numbers should be equal to three: %v\n", nums)
    }

    (*selectedSlice) = append((*selectedSlice), nums)
  }

  lowestLocation := -1

  for i := 0; i < len(initialSeeds); i += 2 {
    fmt.Println(initialSeeds[i])
    for j := initialSeeds[i]; j <= initialSeeds[i] + initialSeeds[i + 1]; j++ {
      soil := convertValue(seedToSoil, j)
      fertilizer := convertValue(soilToFertilizer, soil)
      water := convertValue(fertilizerToWater, fertilizer)
      light := convertValue(waterToLight, water)
      temp := convertValue(lightToTemp, light)
      humidity := convertValue(tempToHumidity, temp)
      location := convertValue(humidityToLocation, humidity)

      if lowestLocation == -1 {
	lowestLocation = location
      }
      lowestLocation = integer.Min(lowestLocation, location)
    }
  }

  return lowestLocation
}

func parseInitialSeeds(line string) []int {
  splitLine := strings.Split(line, ":")
  if len(splitLine) != 2 {
    log.Fatalln("Invalid Inital Seeds Line: " + line)
  }

  return integer.ConvertToIntSlice(splitLine[1], " ")
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
