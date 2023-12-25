package daysixteen

import (
	"errors"
)

type direction int

const (
	LEFT direction = iota
	RIGHT
	UP
	DOWN
)

type tileStack struct {
	tiles []*tile
}

func (s *tileStack) Push(t *tile) {
	s.tiles = append(s.tiles, t)
}

func (s *tileStack) Pop() (*tile, error) {
	if len(s.tiles) == 0 {
		return nil, errors.New("Tile Stack is empty.")
	}

	lastIndex := len(s.tiles) - 1

	t := s.tiles[lastIndex]
	s.tiles = s.tiles[:lastIndex]

	return t, nil
}

func (s *tileStack) Peek() (*tile, error) {
	if len(s.tiles) == 0 {
		return nil, errors.New("Tile Stack is empty.")
	}

	return s.tiles[len(s.tiles)-1], nil
}

func (s *tileStack) IsEmpty() bool {
	return len(s.tiles) == 0
}

type tile struct {
	loc tileLocation
	dir direction
}

type tileLocation struct {
	row int
	col int
}

func NewTile(r, c int, d direction) *tile {
	return &tile{
		loc: tileLocation{
			row: r,
			col: c,
		},
		dir: d,
	}
}

func (t *tile) Row() int {
	return t.loc.row
}

func (t *tile) Col() int {
	return t.loc.col
}

func (t *tile) Loc() tileLocation {
	return t.loc
}

func (t *tile) RowCol() (int, int) {
	return t.loc.row, t.loc.col
}

func (t *tile) Dir() direction {
	return t.dir
}

func NumEnergizedTiles(lines []string) int {
	// row -> col
	visitedTiles := map[tile]bool{}
	visitedCount := map[tileLocation]int{}

	numRows := len(lines)
	numCols := len(lines[0])
	tiles := make([][]byte, numRows)

	for i, l := range lines {
		tiles[i] = []byte(l)
	}

	stack := &tileStack{}

	stack.Push(NewTile(0, 0, RIGHT))

	for !stack.IsEmpty() {
		currTile, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		row := currTile.Row()
		col := currTile.Col()
		loc := currTile.Loc()
		dir := currTile.Dir()

		// Mark it as visited
		visitedTiles[*currTile] = true
		// Increment visited count
		_, ok := visitedCount[loc]
		if !ok {
			visitedCount[loc] = 0
		}
		visitedCount[loc]++

		next := findNextDirection(tiles[row][col], dir)

		for _, n := range next {
			nextRow, nextCol := findNextTileLocation(row, col, n)

			if nextRow < 0 || nextRow > numRows-1 || nextCol < 0 || nextCol > numCols-1 {
				// Out of Bounds
				continue
			}

			nextTile := NewTile(nextRow, nextCol, n)
			_, contains := visitedTiles[*nextTile]
			if contains {
				continue
			}
			stack.Push(nextTile)
		}
	}

	return len(visitedCount)
}

var nextDirections = map[byte]map[direction][]direction{
	byte('.'): {
		UP:    []direction{UP},
		DOWN:  []direction{DOWN},
		LEFT:  []direction{LEFT},
		RIGHT: []direction{RIGHT},
	},
	byte('/'): {
		UP:    []direction{RIGHT},
		DOWN:  []direction{LEFT},
		LEFT:  []direction{DOWN},
		RIGHT: []direction{UP},
	},
	byte('\\'): {
		UP:    []direction{LEFT},
		DOWN:  []direction{RIGHT},
		LEFT:  []direction{UP},
		RIGHT: []direction{DOWN},
	},
	byte('|'): {
		UP:    []direction{UP},
		DOWN:  []direction{DOWN},
		LEFT:  []direction{UP, DOWN},
		RIGHT: []direction{UP, DOWN},
	},
	byte('-'): {
		UP:    []direction{LEFT, RIGHT},
		DOWN:  []direction{LEFT, RIGHT},
		LEFT:  []direction{LEFT},
		RIGHT: []direction{RIGHT},
	},
}

func findNextDirection(c byte, d direction) []direction {
	return nextDirections[c][d]
}

func findNextTileLocation(row, col int, d direction) (int, int) {
	switch d {
	case LEFT:
		return row, col - 1
	case RIGHT:
		return row, col + 1
	case UP:
		return row - 1, col
	case DOWN:
		return row + 1, col
	default:
		return row, col
	}
}
