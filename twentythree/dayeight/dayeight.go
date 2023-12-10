package dayeight

import (
	"log"
	"strings"
)

type stringNode struct {
  Left string
  Right string
}

type node struct {
  Value string
  Type nodeType
  Left string
  Right string
}

type nodeType int

const (
  START_NODE nodeType = 0
  INTERMEDIATE_NODE nodeType = 1
  END_NODE nodeType = 2
)

func GetStepsRequiredToEnd(lines []string) int {
  // First line contains the directions
  directions := []rune(lines[0])
  nodeNetwork := map[string]stringNode{}

  // Second line is empty, so start from third line
  for i := 2; i < len(lines); i++ {
    line := lines[i]
    split := strings.Split(line, " = ")
    if len(split) != 2 {
      log.Fatalf("Network entry are not valid: %v", split[0])
    }

    // (A, B) convert to list of A and B
    nextNodes := strings.Split(strings.TrimRight(strings.TrimLeft(split[1], "("), ")"), ", ")
    if len(nextNodes) != 2 {
      log.Fatalf("Nodes are not valid: %v", split[1])
    }

    currentNode := split[0]
    nodeNetwork[currentNode] = stringNode{nextNodes[0], nextNodes[1]}
  }

  numSteps := 0

  startNode := "AAA"
  currNode := startNode

  for currNode != "ZZZ" {
    // Convert linear array to circular using modular
    directionIndex := numSteps % len(directions)
    if directions[directionIndex] == rune('L') {
      currNode = nodeNetwork[currNode].Left
    } else {
      currNode = nodeNetwork[currNode].Right
    }
    numSteps++
  }

  return numSteps
}

func GetStepsRequiredToZNodes(lines []string) int {
  // First line contains the directions
  directions := []rune(lines[0])
  nodeNetwork := map[string]node{}

  // Second line is empty, so start from third line
  initialNodes := []string{}
  for i := 2; i < len(lines); i++ {
    line := lines[i]
    split := strings.Split(line, " = ")
    if len(split) != 2 {
      log.Fatalf("Network entry are not valid: %v", split[0])
    }

    // (A, B) convert to list of A and B
    nextNodes := strings.Split(strings.TrimRight(strings.TrimLeft(split[1], "("), ")"), ", ")
    if len(nextNodes) != 2 {
      log.Fatalf("Nodes are not valid: %v", split[1])
    }

    nodeValue := split[0]
    nodeType := INTERMEDIATE_NODE

    suffix := nodeValue[len(nodeValue) - 1]

    if suffix == byte('A') {
      nodeType = START_NODE
    } else if suffix == byte('Z') {
      nodeType = END_NODE
    }

    nodeNetwork[nodeValue] = node{Value: nodeValue, Type: nodeType, Left: nextNodes[0], Right: nextNodes[1]}

    if nodeType == START_NODE {
      initialNodes = append(initialNodes, nodeValue)
    }
  }

  // Start from inital nodes
  currNodes := make([]string, len(initialNodes))
  copy(currNodes, initialNodes)

  numSteps := 0
  endSteps := []int{}
  for len(endSteps) != len(initialNodes) {
    // Convert linear array to circular using modular
    directionIndex := numSteps % len(directions)
    shouldMoveLeft := directions[directionIndex] == rune('L')
    numSteps++

    for i, currNodeKey := range currNodes {
      currNode := nodeNetwork[currNodeKey]

      if currNode.Type == END_NODE {
        endSteps = append(endSteps, numSteps - 1)
      }

      if shouldMoveLeft {
        currNodes[i] = currNode.Left
      } else {
        currNodes[i] = currNode.Right
      }
    }
  }

  // Use LCM because of cyclic nature of the end nodes
  return lcm(endSteps[0], endSteps[1], endSteps[2:]...) 
}

func gcd(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
  result := a * b / gcd(a, b)

  for i := 0; i < len(integers); i++ {
    result = lcm(result, integers[i])
  }

  return result
}
