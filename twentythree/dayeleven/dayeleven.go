package dayeleven

import (
	"github.com/amrittb/adventofcode/integer"
)

type coordinate struct {
  X int
  Y int
}

func SumOfShortestDistanceBetweenGalaxies(cosmicMap []string) int {
  return sumOfShortestDistanceBetweenGalaxies(cosmicMap, 2)
}

func SumOfShortestDistanceBetweenOlderGalaxies(cosmicMap []string) int {
  return sumOfShortestDistanceBetweenGalaxies(cosmicMap, 1_000_000)
}

func sumOfShortestDistanceBetweenGalaxies(cosmicMap []string, expansionRate int) int {
  // - Find free x's and y's
  // - Find galaxies and their expanded positions
  // - For each pair find their minimum distance & calculate the sum
  galaxy := byte('#')
  lenY := len(cosmicMap)
  lenX := len(cosmicMap[0])

  galaxyXPositions := integer.IntSet{}
  galaxyYPositions := integer.IntSet{}

  for y := 0; y < lenY; y++ {
    for x := 0; x < lenX; x++ {
      if cosmicMap[y][x] == galaxy {
        galaxyXPositions.Add(x)
        galaxyYPositions.Add(y)
      }
    }
  }

  newGalaxyCords := []coordinate{}

  freeYCount := 0
  for y := 0; y < lenY; y++ {
    if !galaxyYPositions.Contains(y) {
      freeYCount++
      continue
    }

    freeXCount := 0
    for x := 0; x < lenX; x++ {
      if !galaxyXPositions.Contains(x) {
        freeXCount++
        continue
      }

      mapVal := cosmicMap[y][x]

      if mapVal == galaxy {
        newX := x + freeXCount * (expansionRate - 1)
        newY := y + freeYCount * (expansionRate - 1)
        newGalaxyCords = append(newGalaxyCords, coordinate{X: newX, Y: newY})
      }
    }
  }

  sum := 0
  for i := 0; i < len(newGalaxyCords) - 1; i++ {
    for j := i + 1; j < len(newGalaxyCords); j++ {
      sum += shortestPathDistance(newGalaxyCords[i], newGalaxyCords[j])
    }
  }
  
  return sum
}

func shortestPathDistance(c1, c2 coordinate) int {
  return abs(c2.X - c1.X) + abs(c2.Y - c1.Y)
}

func abs(i int) int {
  if i < 0 {
    return i * -1
  }
  
  return i
}

