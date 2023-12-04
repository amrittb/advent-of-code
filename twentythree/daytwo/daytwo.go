package daytwo

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type cubeSelection struct {
	BlueCount  int
	RedCount   int
	GreenCount int
}

func parseGameId(gameIdData string) int {
	gameIdPattern := regexp.MustCompile(`Game\s(\d+)`)
	matches := gameIdPattern.FindStringSubmatch(gameIdData)
	if matches == nil || len(matches) != 2 {
		log.Fatalln("Matches not found")
	}

	// Extract gameId
	gameId, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}

	return gameId
}

func parseCubeSelection(cubeSelectionData string) []cubeSelection {
	selections := strings.Split(cubeSelectionData, ";")
	parsedList := make([]cubeSelection, 0)

	for _, selection := range selections {
		cubes := strings.Split(selection, ",")

		parsedCube := cubeSelection{}

		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")

			res := strings.Split(cube, " ")
			if len(res) != 2 {
				log.Fatalln("Invalid cube: " + cube)
			}

			val, err := strconv.Atoi(res[0])
			if err != nil {
				log.Fatalln(err)
			}

			switch res[1] {
			case "blue":
				parsedCube.BlueCount = val
			case "red":
				parsedCube.RedCount = val
			case "green":
				parsedCube.GreenCount = val
			}
		}

		parsedList = append(parsedList, parsedCube)
	}

	return parsedList
}

// Returns GameId and cube selections
func parseGameData(gameData string) (int, []cubeSelection) {
	res := strings.Split(gameData, ":")
	if len(res) != 2 {
		log.Fatalln("Input data invalid: " + gameData)
	}

	return parseGameId(res[0]), parseCubeSelection(res[1])
}

func SumPossibleGameIds(games []string) int {
	maxRedCount := 12
	maxGreenCount := 13
	maxBlueCount := 14

	sum := 0

	for _, game := range games {
		gameId, draws := parseGameData(game)

		isGamePossible := true
		for _, draw := range draws {
			if draw.BlueCount > maxBlueCount || draw.GreenCount > maxGreenCount || draw.RedCount > maxRedCount {
				isGamePossible = false
				break
			}
		}

		if isGamePossible {
			sum += gameId
		}
	}

	return sum
}

func SumOfPowerOfMinCubes(games []string) int {
	sum := 0

	for _, game := range games {
		_, draws := parseGameData(game)

		minBlueRequired := 0
		minRedRequired := 0
		minGreenRequired := 0

		for _, draw := range draws {
			if draw.BlueCount > minBlueRequired {
				minBlueRequired = draw.BlueCount
			}

			if draw.RedCount > minRedRequired {
				minRedRequired = draw.RedCount
			}

			if draw.GreenCount > minGreenRequired {
				minGreenRequired = draw.GreenCount
			}
		}

		// Add power to the sum
		sum += (minBlueRequired * minRedRequired * minGreenRequired)
	}

	return sum
}
