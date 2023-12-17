package dayten

import (
	"fmt"
)

type pipeType rune

const (
  NONE pipeType = '.'
  START pipeType = 'S'
  NORTH_SOUTH pipeType = '|'
  EAST_WEST pipeType = '-'
  NORTH_EAST pipeType = 'L'
  NORTH_WEST pipeType = 'J'
  SOUTH_EAST pipeType = 'F'
  SOUTH_WEST pipeType = '7'
)

type movement rune

const (
  NORTH movement = 'N'
  SOUTH movement = 'S'
  WEST movement = 'W'
  EAST movement = 'E'
)


// last movement => map of pipeType to next movement
var NEXT_MOVE_MAP_NEW = map[movement]map[pipeType]movement {
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



// last movement => map of pipeType to next movement
var NEXT_MOVE_MAP = map[movement]map[pipeType]movement {
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

  currRow := -1
  currCol := -1
  var lastMove movement

  possibleMoves := []movement{NORTH, SOUTH, EAST, WEST}

  for _, move := range possibleMoves {
    r, c := getNextPosition(startRow, startCol, move)
    nextPipe := pipes[r][c]
    _, isMoveValid := NEXT_MOVE_MAP[move][nextPipe]; if isMoveValid {
      currRow = r
      currCol = c
      lastMove = move
      break
    }
  }

  // Loop until the two positions merge together
  numSteps := 1
  for !(currRow == startRow && currCol == startCol) {
    lastMove = NEXT_MOVE_MAP[lastMove][pipes[currRow][currCol]]
    currRow, currCol = getNextPosition(currRow, currCol, lastMove)
    numSteps++
  }
  return numSteps / 2
}

func NumOfTilesEnclosed(lines []string) int {
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

  currRow := -1
  currCol := -1
  var lastMove movement

  possibleMoves := []movement{NORTH, SOUTH, EAST, WEST}

  for _, move := range possibleMoves {
    r, c := getNextPosition(startRow, startCol, move)
    nextPipe := pipes[r][c]
    _, isMoveValid := NEXT_MOVE_MAP[move][nextPipe]; if isMoveValid {
      currRow = r
      currCol = c
      lastMove = move
      break
    }
  }

  startMove := lastMove

  // Loop until the two positions merge together
  loopTiles := makeBoolMatrix(numRow, numCol)

  loopTiles[startRow][startCol] = true
  for !(currRow == startRow && currCol == startCol) {
    loopTiles[currRow][currCol] = true
    lastMove = NEXT_MOVE_MAP[lastMove][pipes[currRow][currCol]]
    currRow, currCol = getNextPosition(currRow, currCol, lastMove)
  }
  
  startTilePipe := NONE
  reversedLastMove := reverseMove(lastMove)
  for k, v := range NEXT_MOVE_MAP[reverseMove(startMove)] {
    if v == reversedLastMove {
      startTilePipe = k
      break
    }
  }
    
  numEnclosed := 0
  for r, row := range loopTiles {
    pipesSeen := 0
    lastPipe := NONE

    for c, isTileInLoop := range row {
      if !isTileInLoop {
        // When odd number of pipes seen, then mark it as enclosed.
        if pipesSeen % 2 == 1 {
          numEnclosed++
        }
        continue
      }
 
      pipe := pipes[r][c]
      if pipe == START {
        pipe = startTilePipe
      }

      if pipe == NONE || pipe == EAST_WEST {
        continue
      }

      if (pipe == SOUTH_WEST && lastPipe == NORTH_EAST) || (pipe == NORTH_WEST && lastPipe == SOUTH_EAST) {
        // Skip as it only needs to be incremented once
        lastPipe = pipe
        continue
      }

      pipesSeen++
      lastPipe = pipe
    }
  } 

  return numEnclosed
}

func reverseMove(move movement) movement {
  switch move {
  case NORTH:
    return SOUTH
  case SOUTH:
    return NORTH
  case EAST:
    return WEST
  case WEST:
    return EAST
  }
  // Should never happen
  return NORTH
}

func makeBoolMatrix(r, c int) [][]bool {
  matrix := make([][]bool, r)
  for i := range matrix {
    matrix[i] = make([]bool, c)
  }
  return matrix
}

func printBoolMatrixAsBinary(matrix [][]bool) {
  for _, row := range matrix {
    for _, val := range row {
      binVal := 0
      if val {
        binVal = 1
      }
      fmt.Printf("%v ", binVal)
    }
    fmt.Println("")
  }
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
