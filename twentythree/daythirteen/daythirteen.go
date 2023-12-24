package daythirteen

import "fmt"

func SummarizeReflectionPatterns(lines []string) int {
  const rock = rune('#')

  horizontal := []int{}
  vertical := []int{}

  summary := 0

  // Make sure lines end on ""
  if lines[len(lines) - 1] != "" {
    lines = append(lines, "")
  }
  
  i := 0
  for _, l := range lines {
    if l == "" {
      // Calculate here
      vRefIndex := findReflectionIndex(vertical)
      hRefIndex := findReflectionIndex(horizontal)

      summary += vRefIndex + (100 * hRefIndex)

      // Reset values for next block
      horizontal = []int{}
      vertical = []int{}
      i = 0
      continue
    }

    if len(vertical) == 0 {
      vertical = make([]int, len(l))
    }

    horizontal = append(horizontal, 0)
    for j, v := range l {
      if v != rock {
        continue
      }

      horizontal[i] |= 1 << j
      vertical[j] |= 1 << i
    }
    i++
  }

  return summary
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
      // Loop over the matrix & try to fix one cell
      originalHMirror := findReflectionIndexForByteArray(horizontal)
      originalVMirror := findReflectionIndexForByteArray(vertical)

      isMirrorFound := false
      for x := 0; x < len(vertical); x++ {
        for y := 0; y < len(horizontal); y++ {
          original := horizontal[y][x]

          if original == byte('#') {
           horizontal[y][x] = byte('.') 
          } else {
           horizontal[y][x] = byte('#') 
          }

          hRefIndex := findReflectionIndexForByteArray(horizontal)
          horizontal[y][x] = original

          if hRefIndex != 0 && hRefIndex != originalHMirror {
            summary += 100 * hRefIndex
            isMirrorFound = true
            break
          }

          if original == byte('#') {
           vertical[x][y] = byte('.') 
          } else {
           vertical[x][y] = byte('#') 
          }

          vRefIndex := findReflectionIndexForByteArray(vertical)
          vertical[x][y] = original

          if vRefIndex != 0 && vRefIndex != originalVMirror {
            summary += vRefIndex
            isMirrorFound = true
            break
          }
        }

        if isMirrorFound {
          break
        }
      }

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

func findReflectionIndex(values []int) int {
  for i := 0; i < len(values) - 1; i++ {
    isMirrored := false
    if values[i] == values[i+1] {
      // Fan out
      j := i
      k := i + 1
      isMirrored = true
      for j >= 0 && k <= len(values) - 1 {
        if values[j] != values[k] {
          isMirrored = false
          break
        }
        j--
        k++
      }
    }

    if isMirrored {
      return i + 1
    }
  }

  return 0
}

func findReflectionIndexForByteArray(values [][]byte) int {
/*   for _, l := range values {
    fmt.Println(string(l))
  }
 fmt.Println("")
*/  
  for i := 0; i < len(values) - 1; i++ {
    isMirrored := false
    if string(values[i]) == string(values[i+1]) {
      // Fan out
      j := i
      k := i + 1
      isMirrored = true
      for j >= 0 && k <= len(values) - 1 {
        if string(values[j]) != string(values[k]) {
          isMirrored = false
          break
        }
        j--
        k++
      }
    }

    if isMirrored {
      return i + 1
    }
  }

  return 0
}

func reverseBitInt(num, pos int) int {
  return int(reverseBit(uint32(num), uint32(pos)))
}

func reverseBit(num, pos uint32) uint32 {
	var mask uint32 = (1 << pos)
	return (^num & mask) | (num & ^mask)
}

func printBinary(n uint32) {
	fmt.Printf("%v : %032b\n", n, n)
}

