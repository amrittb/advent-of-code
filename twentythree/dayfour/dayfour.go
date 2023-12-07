package dayfour

import (
	"log"
	"strconv"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

func SumOfPointsOfScratchCard(lines []string) int {
  totalPoints := 0

  for _, line := range lines {
    // Remove the first part
    splitLine := strings.Split(line, ":")
    if len(splitLine) != 2 {
      log.Fatalln("Line is invalid")
    }

    nums := strings.Split(splitLine[1], "|")
    if len(nums) != 2 {
      log.Fatalln("Number format is invalid")
    }
    
    winningNumStr, scratchNumStr := nums[0], nums[1]

    winningNums := make(integer.IntSet)
    // Extract winning numbers from first part
    for _, num := range strings.Split(winningNumStr, " ") {
      numAsStr := strings.Trim(num, " ")
      if numAsStr == "" {
        continue
      }

      val, err := strconv.Atoi(numAsStr)
      if err != nil {
        log.Fatalln(err) 
      }

      winningNums.Add(val)
    }

    power := -1
    for _, num := range strings.Split(scratchNumStr, " ") {
      numAsStr := strings.Trim(num, " ")
      if numAsStr == "" {
        continue
      }

      val, err := strconv.Atoi(numAsStr)
      if err != nil {
        log.Fatalln(err) 
      }
      
      if winningNums.Contains(val) {
        power += 1
      }
    }

    if power != -1 {
      totalPoints += integer.Pow(2, power)
    }
  }

  return totalPoints
}

func SumOfScratchCards(lines []string) int {
  scratchCards := make([]int, len(lines))

  for row, line := range lines {
    // Add for original scratch card
    scratchCards[row] += 1
    // Remove the first part
    splitLine := strings.Split(line, ":")
    if len(splitLine) != 2 {
      log.Fatalln("Line is invalid")
    }

    nums := strings.Split(splitLine[1], "|")
    if len(nums) != 2 {
      log.Fatalln("Number format is invalid")
    }

    winningNumStr, scratchNumStr := nums[0], nums[1]

    winningNums := make(integer.IntSet)
    // Extract winning numbers from first part
    for _, num := range strings.Split(winningNumStr, " ") {
      numAsStr := strings.Trim(num, " ")
      if numAsStr == "" {
        continue
      }

      val, err := strconv.Atoi(numAsStr)
      if err != nil {
        log.Fatalln(err)
      }

      winningNums.Add(val)
    }

    power := 0
    for _, num := range strings.Split(scratchNumStr, " ") {
      numAsStr := strings.Trim(num, " ")
      if numAsStr == "" {
        continue
      }

      val, err := strconv.Atoi(numAsStr)
      if err != nil {
        log.Fatalln(err)
      }

      if winningNums.Contains(val) {
        power += 1
      }
    }

    if power != 0 {
      for i := row + 1; i <= min(row + power, len(lines)); i++ {
        scratchCards[i] += scratchCards[row]
      }
    }
  }

  totalScratchCards := 0

  for _, count := range scratchCards {
    // Copy count and one original
    totalScratchCards += count
  }

  return totalScratchCards
}
