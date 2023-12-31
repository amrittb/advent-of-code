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
	actual := SumOfPossibleGameIds(games)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}
}

func Test_SumOfPowerOfMinCubes(t *testing.T) {
	games := []string{
		"Game 1: 5 blue, 2 red; 5 green, 1 blue",
		"Game 2: 1 blue, 5 red, 6 green; 8 red",
	}

	expected := (5 * 2 * 5) + (1 * 8 * 6)
	actual := SumOfPowerOfMinCubes(games)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}

}
