package daythree

import (
	"unicode"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
				prevRow := max(row - 1, 0)
				nextRow := min(row + 1, numRow)
				prevCol := max(col - 1, 0)
				nextCol := min(col + 1, numCol)

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
