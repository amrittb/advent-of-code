package daythree

import (
	"unicode"

	"github.com/amrittb/adventofcode/integer"
)

func SumOfPartNumsOfEngineSchemantic(lines []string) int {
	numRow := len(lines)
	numCol := len(lines[0]) // The column size is always same for ALL rows

	// Create a 2D array of valid positions
	validPos := make([][]bool, numRow)
	for i := range validPos {
		validPos[i] = make([]bool, numCol)
	}

	// In the first loop, 
	// find valid symbol position and their immediate boundary
	for row, line := range lines {
		lineRunes := []rune(line)

		for col, val := range lineRunes {
			if ! unicode.IsDigit(val) && val != '.' {
				// Not a valid symbol, so populate the validPos array
				prevRow := integer.Max(row - 1, 0)
				nextRow := integer.Min(row + 1, numRow)
				prevCol := integer.Max(col - 1, 0)
				nextCol := integer.Min(col + 1, numCol)

				// Previous row
				validPos[prevRow][prevCol] = true
				validPos[prevRow][col] = true
				validPos[prevRow][nextCol] = true

				// Current row
				validPos[row][prevCol] = true
				validPos[row][col] = true
				validPos[row][nextCol] = true

				// Next row
				validPos[nextRow][prevCol] = true
				validPos[nextRow][col] = true
				validPos[nextRow][nextCol] = true
			}	
		}
	}

	sum := 0
	for row, line := range lines {
		lineRunes := []rune(line)

		tmpNum := 0
		isTmpNumValid := false
		for col, val := range lineRunes {
			if unicode.IsDigit(val) {
				currNum := int(val - '0')
				tmpNum = tmpNum * 10 + currNum	
				isTmpNumValid = isTmpNumValid || validPos[row][col]
			} else {
				if tmpNum != 0 && isTmpNumValid {
					sum += tmpNum
				}
				tmpNum = 0
				isTmpNumValid = false
			}
		}

		// Case when the number ends the column
		if tmpNum != 0 && isTmpNumValid {
			sum += tmpNum
		}
	}

	return sum
}

func SumOfGearRatios(lines []string) int {
	numRow := len(lines)
	numCol := len(lines[0]) // The column size is always same for ALL rows

	numMatrix := make([][]*int, numRow)
	for row := range numMatrix {
		numMatrix[row] = make([]*int, numCol)
	}

	// Create 2D array with numbers
	for row, line := range lines {
		lineRunes := []rune(line)

		var tmpNum *int
		numScanned := false
		for col, val := range lineRunes {
			if unicode.IsDigit(val) {
				if tmpNum == nil {
					tmpNum = new(int)
				}

				*tmpNum = *tmpNum * 10 + int(val - '0')

				numMatrix[row][col] = tmpNum
				numScanned = true
			} else {
				if numScanned {
					tmpNum =  nil
				}
				numScanned = false
			}
		}
	}

	sum := 0
	for row, line := range lines {
		lineRunes := []rune(line)

		for col, val := range lineRunes {
			if val == '*' {
				// Gear found, so populate the find adjacent arrays
				prevRow := integer.Max(row - 1, 0)
				nextRow := integer.Min(row + 1, numRow)
				prevCol := integer.Max(col - 1, 0)
				nextCol := integer.Min(col + 1, numCol)

				adjacentNums := make(integer.IntSet)

				// Previous row
				adjacentNums.AddIfNotNil(numMatrix[prevRow][prevCol])
				adjacentNums.AddIfNotNil(numMatrix[prevRow][col])
				adjacentNums.AddIfNotNil(numMatrix[prevRow][nextCol])

				// Current row
				adjacentNums.AddIfNotNil(numMatrix[row][prevCol])
				adjacentNums.AddIfNotNil(numMatrix[row][col])
				adjacentNums.AddIfNotNil(numMatrix[row][nextCol])

				// Next row
				adjacentNums.AddIfNotNil(numMatrix[nextRow][prevCol])
				adjacentNums.AddIfNotNil(numMatrix[nextRow][col])
				adjacentNums.AddIfNotNil(numMatrix[nextRow][nextCol])

				if len(adjacentNums) == 2 {
					gearRatio := 1

					for num := range adjacentNums {
						gearRatio *= num
					}

					sum += gearRatio
				}
			}
		}
	}

	return sum
}

