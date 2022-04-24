package examplefmt

import (
	"bytes"
	"testing"
)

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

func TestPrintConsoleAndSaveToBuffer(t *testing.T) {
	word := "test"
	want := "Printing without formatting: test"
	var output bytes.Buffer
	PrintConsoleAndSaveToBuffer(&output, word);
	if got := output.String(); got != want {
		t.Errorf("PrintConsoleAndSaveToBuffer(os.Stdout, test) = %q, want %q", got, want)
	}
}

func TestPrintfConsoleAndSaveToBuffer(t *testing.T) {
	word := "test"
	want := `Printing with formatting: 
	test`
	var output1 bytes.Buffer
	PrintfConsoleAndSaveToBuffer(&output1, word);
	if got := output1.String(); got != want {
		t.Errorf("PrintfConsoleAndSaveToBuffer(os.Stdout, test) = %q, want %q", got, want)
	}
}
