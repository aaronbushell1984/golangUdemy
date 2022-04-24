package examplefmt

import "testing"

func TestGetBinaryFromInt(t *testing.T) {
	number := 42
	want := "101010"
	if got := GetBinaryFromInt(number); got != want {
		t.Errorf("GetGetBinaryFromInt(number) = %q, want %q", got, want)
	}
}
