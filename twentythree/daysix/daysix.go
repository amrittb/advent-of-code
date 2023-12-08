package daysix

import (
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

func GetWinningCombinations(lines []string) int {
  times := integer.ConvertToIntSlice(strings.Split(lines[0], ":")[1], " ")
  distances := integer.ConvertToIntSlice(strings.Split(lines[1], ":")[1], " ")

  combinations := 1
  for i := 0; i < len(times); i++ {
    winnable := 0
    for t := 1; t < times[i]; t++ {
      d := t * (times[i] - t)
      if d > distances[i] {
	winnable++
      }
    }

    combinations *= winnable
  }

  return combinations
}


func GetWinningCombinationsWithFixedKerning(lines []string) int {
  timeStr := strings.Split(lines[0], ":")[1]
  distanceStr := strings.Split(lines[1], ":")[1]

  time := integer.TrimAndAtoi(timeStr)
  distance := integer.TrimAndAtoi(distanceStr)

  combinations := 0
  for t := 1; t < time; t++ {
    d := t * (time - t)
    if d > distance {
      combinations++
    }
  }

  return combinations
}
