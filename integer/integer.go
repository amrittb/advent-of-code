package integer

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
