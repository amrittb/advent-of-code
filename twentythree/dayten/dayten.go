package dayten

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

  currRow := -1
  currCol := -1
  var lastMove movement

  possibleMoves := []movement{NORTH, SOUTH, EAST, WEST}

  for _, move := range possibleMoves {
    r, c := getNextPosition(startRow, startCol, move)
    nextPipe := pipes[r][c]
    _, isMoveValid := nextMoveMap[move][nextPipe]; if isMoveValid {
      currRow = r
      currCol = c
      lastMove = move
      break
    }
  }

  // Loop until the two positions merge together
  numSteps := 1
  for !(currRow == startRow && currCol == startCol) {
    lastMove = nextMoveMap[lastMove][pipes[currRow][currCol]]
    currRow, currCol = getNextPosition(currRow, currCol, lastMove)
    numSteps++
  }
  return numSteps / 2
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
