package printer

import "fmt"

func PrintByteMatrix(matrix [][]byte) {
  for _, row := range matrix {
    fmt.Println(string(row))
  }
}

func PrintBoolMatrixAsBinary(matrix [][]bool) {
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
