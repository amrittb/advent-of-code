package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func panicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sum := sumCalibrationValues("01-input.txt")
	fmt.Printf("Sum of calibration values: %v\n", sum)

	sum = sumCalibrationValues("01-test-input.txt")
	fmt.Printf("Sum of calibration values: %v\n", sum)
}

func sumCalibrationValues(fileName string) int {
	file, err := os.Open(fileName)
	panicIfErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for _, line := range lines {
		sum += RecoverCalibrationValue(line)
	}

	return sum
}

func RecoverCalibrationValue(line string) int {
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
	}

	return (lNum * 10) + rNum
}
