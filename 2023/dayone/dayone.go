package dayone

import (
	"strconv"
	"strings"
	"unicode"
)

// Part One
func RecoverCalibrationValue(lines []string) int {
	sum := 0
	for _, line := range lines {
		lineRunes := []rune(line)

		l := 0
		r := len(lineRunes) - 1

		lNum := -1
		rNum := -1

		// Loop until we find both numbers
		for lNum == -1 || rNum == -1 {
			lRune := lineRunes[l]
			rRune := lineRunes[r]

			if unicode.IsDigit(lRune) {
				lNum = int(lRune - '0')
			} else {
				l++
			}

			if unicode.IsDigit(rRune) {
				rNum = int(rRune - '0')
			} else {
				r--
			}

			if l > r {
				break
			}
		}

		sum += (lNum * 10) + rNum
	}

	return sum
}

// Part Two
func RecoverAlphaNumericCalibrationValue(line string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	firstNum := 0
	firstIndex := len(line)

	lastNum := 0
	lastIndex := -1

	for num, numStr := range numbers {
		num = num + 1 // Actual num is index + 1
		numToDecimal := strconv.Itoa(num)

		i := strings.Index(line, numStr)
		if i != -1 && i < firstIndex {
			firstIndex = i
			firstNum = num
		}
		i = strings.Index(line, numToDecimal)
		if i != -1 && i < firstIndex {
			firstIndex = i
			firstNum = num
		}

		j := strings.LastIndex(line, numStr)
		if j != -1 && j > lastIndex {
			lastIndex = j
			lastNum = num
		}
		j = strings.LastIndex(line, numToDecimal)
		if j != -1 && j > lastIndex {
			lastIndex = j
			lastNum = num
		}
	}

	return (firstNum * 10) + lastNum
}
