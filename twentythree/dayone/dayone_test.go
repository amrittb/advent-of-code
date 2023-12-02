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

func Test_RecoverAlphaNumbericCallibrationValue(t *testing.T) {
	lines := []string{
		"eightwo",
		"a14bnineight",
	}

	expected := 82 + 18

	actual := RecoverAlphaNumericCalibrationValue(lines)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}
}
