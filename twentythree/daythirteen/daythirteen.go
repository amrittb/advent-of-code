package daythirteen

import (
	"fmt"

	"github.com/amrittb/adventofcode/integer"
)

func SummarizeReflectionPatterns(lines []string) int {
  horizontal := [][]byte{}
  vertical := [][]byte{}

  summary := 0

  // Make sure lines end on ""
  if lines[len(lines) - 1] != "" {
    lines = append(lines, "")
  }
  
  for _, l := range lines {
    if l == "" {
      // Calculate here
      vRefIndex := findReflectionIndex(vertical)
      hRefIndex := findReflectionIndex(horizontal)

      summary += vRefIndex + (100 * hRefIndex)

      // Reset values for next block
      horizontal = [][]byte{}
      vertical = [][]byte{}
      continue
    }

    if len(vertical) == 0 {
      vertical = make([][]byte, len(l))
    }

    horizontal = append(horizontal, []byte(l))
    for j, v := range l {
      vertical[j] = append(vertical[j], byte(v))
    }
  }

  return summary
}

func findReflectionIndex(values [][]byte) int {
  for i := 1; i < len(values); i++ {
    size := integer.Min(i, len(values) - i)
    above := values[i-size:i]    
    below := values[i:i+size]

    isEqual := false
    for s := 0; s < size; s++ {
      isEqual = string(above[size - s - 1]) == string(below[s])

      if !isEqual {
        break
      }
    }

    if isEqual {
      return i
    }
  }

  return 0
}

func SummarizeSmudgeFixedReflectionPatterns(lines []string) int {
  horizontal := [][]byte{}
  vertical := [][]byte{}

  summary := 0

  // Make sure lines end on ""
  if lines[len(lines) - 1] != "" {
    lines = append(lines, "")
  }
  
  for _, l := range lines {
    if l == "" {
      // fmt.Println("Horizontal")
      summary += 100 * findSmudgeReflectionIndex(horizontal)
      // fmt.Println("Vertical")
      summary += findSmudgeReflectionIndex(vertical)

      // Reset values for next block
      horizontal = [][]byte{}
      vertical = [][]byte{}
      continue
    }

    if len(vertical) == 0 {
      vertical = make([][]byte, len(l))
    }

    horizontal = append(horizontal, []byte(l))
    for j, v := range l {
      vertical[j] = append(vertical[j], byte(v))
    }
  }

  return summary
}

func findSmudgeReflectionIndex(values [][]byte) int {
  for i := 1; i < len(values); i++ {
    size := integer.Min(i, len(values) - i)
    above := values[i-size:i]    
    below := values[i:i+size]

    diff := 0
    for s := 0; s < size; s++ {
      for j, c := range(above[size - s - 1]) {
        if c != below[s][j] {
          diff++
        }
      }
    }

    if diff == 1 {
      return i
    }
  }

  return 0
}

func printByteSlices(val [][]byte) {
  for _, v := range val {
    fmt.Println(string(v))
  }
}
