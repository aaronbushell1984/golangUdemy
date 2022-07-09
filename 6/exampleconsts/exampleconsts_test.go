package exampleconsts_test

import "testing"

func TestReturnTwoPiConsts(t *testing.T) {
	want := "We have returned 3.1416 and 3.1416"
	if got := ReturnTwoFloatVariablesOrConstantsAfterUpdatingOne(); got != want {
		t.Errorf("ReturnTwoFloatVariablesOrConstantsAfterUpdatingOne = %q, want %q", got, want)
	}
}