package main

import "testing"

func Test_RecoverCalibrationValue_MultipleDigit(t *testing.T) {
	line := "a7dasfj0ie08380asjdf0923b"

	expected := 73
	actual := RecoverCalibrationValue(line)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}
}

func Test_RecoverCalibrationValue_SingleDigit(t *testing.T) {
	line := "a7b"

	expected := 77
	actual := RecoverCalibrationValue(line)

	if actual != expected {
		t.Errorf("got %v but wanted %v\n", actual, expected)
	}
}
