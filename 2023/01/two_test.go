package main

import "testing"

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
