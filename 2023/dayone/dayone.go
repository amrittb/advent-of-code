package dayone

import (
	"unicode"
)

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
