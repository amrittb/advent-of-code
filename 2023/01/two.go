package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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
		sum += RecoverAlphaNumericCalibrationValue(line)
	}

	fmt.Printf("Answer: %v\n", sum)
}

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
