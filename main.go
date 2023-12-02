package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	twentythreedayone "github.com/amrittb/adventofcode/twentythree/dayone"
)

type Process func([]string) int

func main() {
	yearArg := flag.Int("year", 2023, "Which year to run?")
	dayArg := flag.Int("day", 0, "Which day to run?")
	partArg := flag.Int("part", 0, "Which part to run?")
	fileNameArg := flag.String("file", "", "Input file")
	flag.Parse()

	year := *yearArg
	day := *dayArg
	part := *partArg
	fileName := *fileNameArg

	validCommands := map[int]map[int]map[int]Process{
		2023: {
			1: {
				1: twentythreedayone.RecoverCalibrationValue,
				2: twentythreedayone.RecoverAlphaNumericCalibrationValue,
			},
		},
	}

	yearMap, yearOk := validCommands[year]
	if !yearOk {
		log.Fatalf("Year argument is invalid: %v\n", year)
	}

	dayMap, dayOk := yearMap[day]
	if !dayOk {
		log.Fatalf("Day argument is invalid: %v\n", day)
	}

	partCommand, partOk := dayMap[part]
	if !partOk {
		log.Fatalf("Part argument is invalid: %v\n", part)
	}

	if fileName == "" {
		log.Fatalf("File argument is invalid: %v\n", fileName)
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if fileInfo.IsDir() {
		log.Fatalf("File %v is a directory, not a file.\n", fileName)
	}

	log.Printf("Running with Year: %v, Day: %v, Part: %v, FileName: %v", year, day, part, fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	ans := partCommand(lines)

	fmt.Printf("Answer: %v\n", ans)
}