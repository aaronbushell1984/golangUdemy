package examplefmt

import "testing"

func TestGetBinaryFromInt(t *testing.T) {
	number := 42
	want := "101010"
	if got := GetBinaryFromInt(number); got != want {
		t.Errorf("GetGetBinaryFromInt(number) = %q, want %q", got, want)
	}
}

func TestGetHexidecimalFromInt(t *testing.T) {
	number := 42
	want := "2a"
	if got := GetHexidecimalFromInt(number); got != want {
		t.Errorf("GetHexidecimalFromInt(number) = %q, want %q", got, want)
	}
}

func TestGetOctalHexidecimalFromInt(t *testing.T) {
	number := 42
	want := "0x2a"
	if got := GetOctalHexidecimalFromInt(number); got != want {
		t.Errorf("GetOctalHexidecimalFromInt(number) = %q, want %q", got, want)
	}
}


