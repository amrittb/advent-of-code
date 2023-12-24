package daythirteen

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

func findReflectionIndexWithPositions(values []int, start, end int) int {
  foundOnce := false

  for start < end {
    if values[start] == values[end] {
      start++
      end--
      foundOnce = true
    } else {
      if foundOnce {
        return -1
      }

      if start < end {
        start++
      } else {
        start--
      }
    }  
  }

  return start
}

