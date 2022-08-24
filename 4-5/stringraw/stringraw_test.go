package stringraw

import "testing"

func TestGetRawStringWithBackticks(t *testing.T) {
	want := `When using backticks "quotes" are printed and
	newlines are interpreted without use of \n`
	if got := GetRawStringWithBackticks(); got != want {
		t.Errorf("GetRawStringWithBackticks() = %q, want %q", got, want)
	}
}
