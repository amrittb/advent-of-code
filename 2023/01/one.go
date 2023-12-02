package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main2() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Invalid args. Please provide input filename.")
	}

	fileName := args[0]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if fileInfo.IsDir() {
		log.Fatalf("File %v is a directory, not a file.\n", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += RecoverCalibrationValue(line)
	}

	fmt.Printf("Answer: %v\n", sum)
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

		if l > r {
			break
		}
	}

	return (lNum * 10) + rNum
}
