package integer

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

type IntSet map[int]bool

func (set *IntSet) Add(val int) {
	(*set)[val] = true
}

func (set *IntSet) AddIfNotNil(val *int) {
	if val != nil {
		(*set).Add(*val)
	}
}

func (set *IntSet) Contains(val int) bool {
	return (*set)[val]
}

// Helper functions
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pow(base, exponent int) int {
	if exponent == 0 {
		return 1
	}

	res := base
	for i := 2; i <= exponent; i++ {
		res *= base
	}
	return res
}

func ConvertToIntSlice(s, sep string) []int {
	ints := []int{}
	intStrs := strings.Split(strings.Trim(s, " "), sep)
	for _, intStr := range intStrs {
		intStr = strings.Trim(intStr, " ")
		if intStr == "" {
			continue
		}
		i, err := strconv.Atoi(intStr)
		if err != nil {
			log.Fatalln(err)
		}
		ints = append(ints, i)
	}

	return ints
}

func TrimAndAtoi(s string) int {
	i := 0
	lineRune := []rune(s)
	for _, r := range lineRune {
		if unicode.IsDigit(r) {
			i = i*10 + int(r-'0')
		}
	}
	return i
}
