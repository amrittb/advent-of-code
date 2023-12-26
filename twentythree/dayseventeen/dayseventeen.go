package dayseventeen

import (
	// "fmt"

	"github.com/amrittb/adventofcode/integer"
)

func LeastHeatLoss(lines []string) int {
	numRows := len(lines)
	// numCols := len(lines[0])

	heatLoss := make([][]int, numRows)

	for i, l := range lines {
		heatLoss[i] = integer.ConvertToIntSlice(l, "")
	}

	return 0
}
