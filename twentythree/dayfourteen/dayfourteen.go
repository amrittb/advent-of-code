package dayfourteen

func TotalLoadOnNorthBeams(lines []string) int {
  roundRock := byte('O')
  cubeRock := byte('#')

  numRows := len(lines)
  numCols := len(lines[0])

  nextRockPositions := make([]int, numCols)

  numRoundRocks := make([]int, numRows)

  for i := 0; i < numRows; i++ {
    for j := 0; j < numCols; j++ {
      currentRock := lines[i][j]
      if currentRock == roundRock {
        numRoundRocks[nextRockPositions[j]]++
        nextRockPositions[j]++
      } else if currentRock == cubeRock {
        nextRockPositions[j] = i + 1
      }
    }
  }

  sum := 0

  for i, v := range numRoundRocks {
    sum += (numRows - i) * v
  }

  return sum
}
