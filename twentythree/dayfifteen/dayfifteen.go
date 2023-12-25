package dayfifteen

import (
	"strings"
)

func SumOfSeqHashes(lines []string) int {
	line := strings.Trim(lines[0], "\n")

	seqSlice := strings.Split(line, ",")

	sum := 0
	for _, s := range seqSlice {
		sum += hash(s)
	}

	return sum
}

func hash(s string) int {
	hash := 0

	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	return hash
}
