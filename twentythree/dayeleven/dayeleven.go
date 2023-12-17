package dayeleven

import (
	"fmt"

	"github.com/amrittb/adventofcode/integer"
)

type coordinate struct {
  X int
  Y int
}

func SumOfShortestDistanceBetweenGalaxies(cosmicMap []string) int {
  // - Find galaxies
  // - Expand the universe
  // - Find galaxies and their positions
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

  // New len = original len + delta len (total len - len of found galaxies in an axis)
  newLenX := lenX + (lenX - len(galaxyXPositions))
  newLenY := lenY + (lenY - len(galaxyYPositions))

  newCosmicMap := make([][]byte, newLenY)

  newY := 0
  for y := 0; y < lenY; y++ {
    newXMap := make([]byte, newLenX)

    newX := 0
    for x := 0; x < lenX; x++ {
      mapVal := cosmicMap[y][x]

      if mapVal == galaxy {
        newGalaxyCords = append(newGalaxyCords, coordinate{X: newX, Y: newY})
      }

      newXMap[newX] = mapVal
      newX++

      // Expand in X-Axis
      if !galaxyXPositions.Contains(x) {
        newXMap[newX] = byte('.')
        newX++
      }
    }

    newCosmicMap[newY] = newXMap
    newY++

    if !galaxyYPositions.Contains(y) {
      // 
      anotherNewXMap := make([]byte, newLenX) 
      copy(anotherNewXMap, newXMap) 

      newCosmicMap[newY] = anotherNewXMap
      newY++
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

func printByteMatrixAsString(input [][]byte) {
  for _, r := range input {
    fmt.Println(string(r))
  }
}

func shortestPathDistance(c1, c2 coordinate) int {
  distance := 0

  x, y := c1.X, c1.Y

  for x != c2.X || y != c2.Y {
    if y < c2.Y {
      y++
      distance++
    } else if y > c2.Y {
      y--
      distance++
    }

    if x < c2.X {
      x++
      distance++
    } else if x > c2.X {
      x--
      distance++
    }
  }

  return distance
}

