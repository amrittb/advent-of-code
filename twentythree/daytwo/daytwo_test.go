package daytwo

import (
	"testing"
)

func Test_SumPossibleGameIds(t *testing.T) {
	games := []string{
		"Game 1: 19 blue, 2 red; 5 green",
		"Game 2: 1 blue, 5 red, 6 green",
	}

	expected := 2
	actual := SumPossibleGameIds(games)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}

}
