package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	day := *flag.Int("day", 0, "Which day to run?")
	part := *flag.Int("part", 0, "Which part to run?")
	fileName := *flag.String("file", "", "Input file")
	flag.Parse()

	if day == 0 {
		log.Fatalf("Day argument is invalid: %v\n", day)
	}

	if part == 0 {
		log.Fatalf("Part argument is invalid: %v\n", part)
	}

	if fileName == "" {
		log.Fatalf("File argument is invalid: %v\n", fileName)
	}

	// dayCommands := make(map[int]Command)

	log.Printf("Day: %v, Part: %v, FileName: %v", day, part, fileName)

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

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	sum := RecoverCalibrationValue(lines)

	fmt.Printf("Answer: %v\n")
}
