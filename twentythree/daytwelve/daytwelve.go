package daytwelve

import (
	"fmt"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

const operational = byte('.')
const damaged = byte('#')
const unknown = byte('?')

func SumOfPossibleArrangements(lines []string) int {
  totalCount := 0
  for _, row := range lines {
    split := strings.Split(row, " ") 
    conditionRecord := []byte(split[0])
    damagedSprings := integer.ConvertToIntSlice(split[1], ",")

    totalCount += count(conditionRecord, damagedSprings)
  }
  return totalCount
}

func SumOfUnfoldedPossibleArrangements(lines []string) int {
  totalCount := 0
  for _, row := range lines {
    split := strings.Split(row, " ") 
    unfoldedRecordSlice := strings.Join([]string{split[0],split[0],split[0],split[0],split[0]}, "?")
    unfoldedDamagedSlice := strings.Join([]string{split[1],split[1],split[1],split[1],split[1]}, ",")
    conditionRecord := []byte(unfoldedRecordSlice)
    damagedSprings := integer.ConvertToIntSlice(unfoldedDamagedSlice, ",")

    totalCount += count(conditionRecord, damagedSprings)
  }
  return totalCount
}

func count(conditionRecord []byte, damagedSprings []int) int {
  // - Find the count & locations of unknown in the conditionRecords
  // - Generate ALL possible values in the unknown
  // - Match with the damagedSprings count and count valid arragements

  unknownPositions := integer.IntSet{}
  for i, r := range conditionRecord {
    if r == unknown {
      unknownPositions.Add(i)
    }
  }

  generatedArrangements := generate(conditionRecord, unknownPositions)

  validCount := 0
  for _, arrangement := range generatedArrangements {
    if isArrangementValid(arrangement, damagedSprings) {
      validCount++
    }
  }

  fmt.Println(validCount)

  return validCount
}

func generate(originalRecord []byte, unknownPositions integer.IntSet) [][]byte {
  // Since the unknown values can be either '.' or '#' (2 possible values),
  // the total no. of possible combinations will be 2 ^ numUnknown
  // which can be representated as left shift of numUnknown-1
  // 2 << (1 - 1) = 2 << 0 = 2
  // 2 << (2 - 1) = 2 << 1 = 4
  originalLen := len(originalRecord)
  unknownLen := len(unknownPositions)

  totalNumCombinations := 2 << (unknownLen - 1)

  result := make([][]byte, totalNumCombinations)

  for i := 0; i < totalNumCombinations; i++ {
    generatedRecord := make([]byte, originalLen)

    uPos := 0
    for j := 0; j < originalLen; j++ {
      if unknownPositions.Contains(j) {
	if IsBitSet(i, uPos) {
	  generatedRecord[j] = damaged
	} else {
	  generatedRecord[j] = operational
	}
	uPos++
      } else {
	generatedRecord[j] = originalRecord[j]
      }
    }

    result[i] = generatedRecord
  }

  return result
}

func IsBitSet(val, pos int) bool {
  return (val & (1 << pos)) != 0
}

func isArrangementValid(testRecord []byte, damagedSprings []int) bool {
  i := 0
  last := unknown
  lastDamagedCount := 0

  // Append operation at the end to ignore checking boundary cases
  testRecord = append(testRecord, operational)

  for _, current := range testRecord {
    if current == damaged {
      lastDamagedCount++
    } else {
      // If operational
      if last == damaged {
	// Switched from damanged to operational
	if i >= len(damagedSprings) {
	  // When a damaged spring is found,
	  // and lookup index is already more than remaining damaged springs,
	  // it's an invalid arragement
	  return false
	}

	if lastDamagedCount != damagedSprings[i] {
	  return false
	}

	i++
      }
      lastDamagedCount = 0
    }

    last = current
  }

  return i == len(damagedSprings)
}

func printByteMatrixAsBinary(matrix [][]byte) {
  for _, row := range matrix {
    fmt.Println(string(row))
  }
}

