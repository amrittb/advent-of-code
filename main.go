package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	twentythreedayeight "github.com/amrittb/adventofcode/twentythree/dayeight"
	twentythreedayeleven "github.com/amrittb/adventofcode/twentythree/dayeleven"
	twentythreedayfifteen "github.com/amrittb/adventofcode/twentythree/dayfifteen"
	twentythreedayfive "github.com/amrittb/adventofcode/twentythree/dayfive"
	twentythreedayfour "github.com/amrittb/adventofcode/twentythree/dayfour"
	twentythreedayfourteen "github.com/amrittb/adventofcode/twentythree/dayfourteen"
	twentythreedaynine "github.com/amrittb/adventofcode/twentythree/daynine"
	twentythreedayone "github.com/amrittb/adventofcode/twentythree/dayone"
	twentythreedayseven "github.com/amrittb/adventofcode/twentythree/dayseven"
	twentythreedaysix "github.com/amrittb/adventofcode/twentythree/daysix"
	twentythreedaysixteen "github.com/amrittb/adventofcode/twentythree/daysixteen"
	twentythreedayten "github.com/amrittb/adventofcode/twentythree/dayten"
	twentythreedaythirteen "github.com/amrittb/adventofcode/twentythree/daythirteen"
	twentythreedaythree "github.com/amrittb/adventofcode/twentythree/daythree"
	twentythreedaytwelve "github.com/amrittb/adventofcode/twentythree/daytwelve"
	twentythreedaytwo "github.com/amrittb/adventofcode/twentythree/daytwo"
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
			2: {
				1: twentythreedaytwo.SumOfPossibleGameIds,
				2: twentythreedaytwo.SumOfPowerOfMinCubes,
			},
			3: {
				1: twentythreedaythree.SumOfPartNumsOfEngineSchemantic,
				2: twentythreedaythree.SumOfGearRatios,
			},
			4: {
				1: twentythreedayfour.SumOfPointsOfScratchCard,
				2: twentythreedayfour.SumOfScratchCards,
			},
			5: {
				1: twentythreedayfive.FindLowestLocation,
				2: twentythreedayfive.FindLowestLocationOfSeedRanges,
			},
			6: {
				1: twentythreedaysix.GetWinningCombinations,
				2: twentythreedaysix.GetWinningCombinationsWithFixedKerning,
			},
			7: {
				1: twentythreedayseven.TotalWinnings,
				2: twentythreedayseven.TotalWinningsWithJoker,
			},
			8: {
				1: twentythreedayeight.GetStepsRequiredToEnd,
				2: twentythreedayeight.GetStepsRequiredToZNodes,
			},
			9: {
				1: twentythreedaynine.SumOfExtrapolatedValues,
				2: twentythreedaynine.SumOfBackwardExtrapolatedValues,
			},
			10: {
				1: twentythreedayten.NumStepsOfFarthestPointInLoop,
				2: twentythreedayten.NumOfTilesEnclosed,
			},
			11: {
				1: twentythreedayeleven.SumOfShortestDistanceBetweenGalaxies,
				2: twentythreedayeleven.SumOfShortestDistanceBetweenOlderGalaxies,
			},
			12: {
				1: twentythreedaytwelve.SumOfPossibleArrangements,
				2: twentythreedaytwelve.SumOfUnfoldedPossibleArrangements,
			},
			13: {
				1: twentythreedaythirteen.SummarizeReflectionPatterns,
				2: twentythreedaythirteen.SummarizeSmudgeFixedReflectionPatterns,
			},
			14: {
				1: twentythreedayfourteen.TotalLoadOnNorthBeams,
				2: twentythreedayfourteen.TotalLoadOnNorthBeamsAfterCycles,
			},
			15: {
				1: twentythreedayfifteen.SumOfSeqHashes,
			},
			16: {
				1: twentythreedaysixteen.NumEnergizedTiles,
				2: twentythreedaysixteen.MaxEnergizedTiles,
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
