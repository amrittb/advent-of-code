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
	line01 := "eightwo"
	line02 := "a14bnineight"

	expected01 := 82
	expected02 := 18

	actual01 := RecoverAlphaNumericCalibrationValue(line01)
	actual02 := RecoverAlphaNumericCalibrationValue(line02)

	if actual01 != expected01 {
		t.Errorf("got %v but wanted %v\n", actual01, expected01)
	}

	if actual02 != expected02 {
		t.Errorf("got %v but wanted %v\n", actual02, expected02)
	}

}
