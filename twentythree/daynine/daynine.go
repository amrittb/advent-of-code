package daynine

import (
	"log"
	"strconv"
	"strings"
)

func SumOfExtrapolatedValues(lines []string) int {
  sum := 0

  for _, line := range lines {
    split := strings.Split(line, " ")
    readings := make([]int, len(split))

    for i, s := range split {
      val, err := strconv.Atoi(s)
      if err != nil {
        log.Fatalf("Reading cannot be converted to slice of int: %v", s)
      }
      readings[i] = val
    }

    sum += extrapolate(readings)
  }

  return sum
}

func SumOfBackwardExtrapolatedValues(lines []string) int {
  sum := 0

  for _, line := range lines {
    split := strings.Split(line, " ")
    readings := make([]int, len(split))

    for i, s := range split {
      val, err := strconv.Atoi(s)
      if err != nil {
        log.Fatalf("Reading cannot be converted to slice of int: %v", s)
      }
      readings[i] = val
    }

    sum += backwardExtrapolate(readings)
  }

  return sum
}

func extrapolate(input []int) int {
  last := input[len(input) - 1]
  diff := make([]int, len(input) - 1)

  zeroCount := 0
  for i := 1; i < len(input); i++ {
    d := input[i] - input[i - 1]
    if d == 0 {
      zeroCount++
    }
    diff[i - 1] = d
  }

  if zeroCount == len(diff) {
    return last
  }

  return last + extrapolate(diff)
}

func backwardExtrapolate(input []int) int {
  first := input[0]
  diff := make([]int, len(input) - 1)

  zeroCount := 0
  for i := 1; i < len(input); i++ {
    d := input[i] - input[i - 1]
    if d == 0 {
      zeroCount++
    }
    diff[i - 1] = d
  }

  if zeroCount == len(diff) {
    return first
  }

  return first - backwardExtrapolate(diff)
}
