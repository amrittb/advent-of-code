package dayten

import (
	// "fmt"
	"log"

	"github.com/amrittb/adventofcode/integer"
)

type pipeType rune

const (
  NONE pipeType = '.'
  START pipeType = 'S'
  NORTH_SOUTH pipeType = '|'
  EAST_WEST pipeType = '-'
  NORTH_EAST pipeType = 'F'
  NORTH_WEST pipeType = '7'
  SOUTH_EAST pipeType = 'L'
  SOUTH_WEST pipeType = 'J'
)

type movement rune

const (
  NORTH movement = 'N'
  SOUTH movement = 'S'
  WEST movement = 'W'
  EAST movement = 'E'
)

func NumStepsOfFarthestPointInLoop(lines []string) int {
  numRow := len(lines)
  numCol := len(lines[0])
  pipes := make([][]pipeType, numRow)

  for i := 0; i < numRow; i ++ {
    pipes[i] = make([]pipeType, numCol)
  }

  startRow := -1
  startCol := -1
  for row, line := range lines {
    lineRunes := []rune(line)

    for col, pipeTypeChar := range lineRunes {
      pipeType := parsePipeTypeChar(pipeTypeChar)
      pipes[row][col] = pipeType

      if pipeType == START {
        startRow = row
	startCol = col
      }
    }
  }

  // last movement => map of pipeType to next movement
  nextMoveMap := map[movement]map[pipeType]movement {
    // to north (allowed) = |, F, 7
    NORTH: {
      NORTH_SOUTH: NORTH,
      SOUTH_EAST: EAST,
      SOUTH_WEST: WEST, 
    },
    // to south (allowed) = |, L, J
    SOUTH: {
      NORTH_SOUTH: SOUTH,
      NORTH_EAST: EAST,
      NORTH_WEST: WEST, 
    },
    // to east (allowed) = -, J, 7
    EAST: {
      EAST_WEST: EAST,
      NORTH_WEST: NORTH,
      SOUTH_WEST: SOUTH, 
    },
    // to west (allowed) = -, L, F
    WEST: {
      EAST_WEST: WEST,
      NORTH_EAST: NORTH,
      SOUTH_EAST: SOUTH, 
    },
  }

  prevRow := integer.Max(startRow - 1, 0)
  nextRow := integer.Min(startRow + 1, numRow - 1)

  prevCol := integer.Max(startCol - 1, 0)
  nextCol := integer.Min(startCol + 1, numCol - 1)

  currPos := [][]int{}

  moves := []movement{}

  northPipe := pipes[prevRow][startCol]
  _, isNorthValid := nextMoveMap[NORTH][northPipe]; if isNorthValid {
    currPos = append(currPos, []int{prevCol, startCol})
    moves = append(moves, NORTH)
  }

  southPipe := pipes[nextRow][startCol]
  _, isSouthValid := nextMoveMap[SOUTH][southPipe]; if isSouthValid {
    currPos = append(currPos, []int{nextRow, startCol})
    moves = append(moves, SOUTH)
  }

  eastPipe := pipes[startRow][nextCol]
  _, isEastValid := nextMoveMap[EAST][eastPipe]; if isEastValid {
    currPos = append(currPos, []int{startRow, nextCol})
    moves = append(moves, EAST)
  }

  westPipe := pipes[startRow][prevCol]
  _, isWestValid := nextMoveMap[WEST][westPipe]; if isWestValid {
    currPos = append(currPos, []int{startRow, prevCol})
    moves = append(moves, WEST)
  }

  if len(moves) != 2 {
    log.Fatalf("Found less or more than 2 possible paths: %v", currPos)
  }

  // Loop until the two positions merge together
  numSteps := 1
  for true {
    if currPos[0][0] == currPos[1][0] && currPos[0][1] == currPos[1][1] {
      break
    }
    for i := 0; i < len(currPos); i++ {
      currRow := currPos[i][0]
      currCol := currPos[i][1]
      moves[i] = nextMoveMap[moves[i]][pipes[currRow][currCol]]
      nextRow, nextCol = getNextPosition(currRow, currCol, moves[i])
      currPos[i][0] = nextRow
      currPos[i][1] = nextCol
    }
    numSteps++
  }
  return numSteps
}

func getNextPosition(r, c int, m movement) (int, int) {
  switch m {
  case NORTH:
    return r - 1, c
  case SOUTH:
    return r + 1, c
  case EAST:
    return r, c + 1
  case WEST:
    return r, c - 1
  default:
    return r, c
  }
}

func parsePipeTypeChar(r rune) pipeType {
  switch r {
  case rune('|'):
    return NORTH_SOUTH
  case rune('-'):
    return EAST_WEST
  case rune('L'):
    return NORTH_EAST
  case rune('J'):
    return NORTH_WEST
  case rune('7'):
    return SOUTH_WEST
  case rune('F'):
    return SOUTH_EAST
  case rune('S'):
    return START
  default:
    return NONE
  }
}
