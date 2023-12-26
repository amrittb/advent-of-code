package dayeighteen

import (
	"strconv"
	"strings"

	"github.com/amrittb/adventofcode/integer"
)

func VolumeOfLava(lines []string) int {
	points := []*integer.IntPair{integer.NewIntPair(0, 0)}

	numBoundaryPoints := 0
	for _, l := range lines {
		split := strings.Split(l, " ")
		dir := split[0]
		steps, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		nextPoint := getNextPoint(points[len(points)-1], dir, steps)
		points = append(points, nextPoint)
		numBoundaryPoints += steps
	}

	return calculateAreaEnclosedByPoints(points, numBoundaryPoints)
}

func VolumeOfLavaFromHex(lines []string) int {
	points := []*integer.IntPair{integer.NewIntPair(0, 0)}

	numBoundaryPoints := 0
	for _, l := range lines {
		split := strings.Split(l, " ")
		dir, steps := extractFromHex(split[2])
		nextPoint := getNextPoint(points[len(points)-1], dir, steps)
		points = append(points, nextPoint)
		numBoundaryPoints += steps
	}

	return calculateAreaEnclosedByPoints(points, numBoundaryPoints)
}

func extractFromHex(hex string) (string, int) {
	hex = strings.Trim(hex, " ")
	stepsRaw := hex[2:7]
	dirRaw := hex[7]

	numberStr := strings.Replace(stepsRaw, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	steps, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		panic(err)
	}

	var dir string
	switch dirRaw {
	case byte('0'):
		dir = "R"
	case byte('1'):
		dir = "D"
	case byte('2'):
		dir = "L"
	case byte('3'):
		dir = "U"
	}

	return dir, int(steps)
}

func getNextPoint(intPair *integer.IntPair, dir string, steps int) *integer.IntPair {
	dx := 0
	dy := 0

	switch dir {
	case "R":
		dx = steps
	case "L":
		dx = -steps
	case "U":
		dy = steps
	case "D":
		dy = -steps
	}

	return integer.NewIntPair(intPair.First()+dx, intPair.Second()+dy)
}

func calculateAreaEnclosedByPoints(points []*integer.IntPair, numBoundaryPoints int) int {
	sum := 0

	// Shoelace formula
	for i := 0; i < len(points)-1; i++ {
		j := i + 1
		a := points[i]
		b := points[j]
		sum += (a.First() * b.Second()) - (a.Second() * b.First())
	}

	sum = integer.Abs(sum)

	// Since area returned also includes the half unit of each boundary point,
	// we need to calculate the correct area
	area := sum / 2

	// Pick's Theorm
	// A = i + (b/2) + 1
	// A = area
	// i = num of interior points
	// b = num of boundary points
	numInnerPoints := area - (numBoundaryPoints / 2) + 1

	// Since each point represent a unit sq area,
	// total area is the sum of num of inner points & num of boundary points
	return numInnerPoints + numBoundaryPoints

}
