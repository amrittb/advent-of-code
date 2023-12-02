package dayone

import "testing"

func Test_RecoverCalibrationValue(t *testing.T) {
	lines := []string{
		"a7dasfj0ie08380asjdf0923b",
		"a7b",
	}

	expected := 73 + 77
	actual := RecoverCalibrationValue(lines)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}
}
