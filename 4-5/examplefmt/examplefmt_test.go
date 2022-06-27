package examplefmt

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetBinaryFromInt(t *testing.T) {
	number := 42
	want := "101010"
	if got := GetBinaryFromInt(number); got != want {
		t.Errorf("GetGetBinaryFromInt(number) = %q, want %q", got, want)
	}
}

func ExampleGetBinaryFromInt() {
	fmt.Println(GetBinaryFromInt(42))
	// Output:
	// 101010
}

func TestGetHexidecimalFromInt(t *testing.T) {
	number := 42
	want := "2a"
	if got := GetHexidecimalFromInt(number); got != want {
		t.Errorf("GetHexidecimalFromInt(number) = %q, want %q", got, want)
	}
}

func ExampleGetHexidecimalFromInt() {
	fmt.Println(GetHexidecimalFromInt(42))
	// Output:
	// 2a
}

func TestGetOctalHexidecimalFromInt(t *testing.T) {
	number := 42
	want := "0x2a"
	if got := GetOctalHexidecimalFromInt(number); got != want {
		t.Errorf("GetOctalHexidecimalFromInt(number) = %q, want %q", got, want)
	}
}

func ExampleGetOctalHexidecimalFromInt() {
	fmt.Println(GetOctalHexidecimalFromInt(42))
	// Output:
	// 0x2a
}

func TestPrintConsoleAndSaveToBuffer(t *testing.T) {
	word := "test"
	want := "Printing without formatting: test"
	var output bytes.Buffer
	PrintConsoleAndSaveToBuffer(&output, word)
	if got := output.String(); got != want {
		t.Errorf("PrintConsoleAndSaveToBuffer(os.Stdout, test) = %q, want %q", got, want)
	}
}

func ExamplePrintConsoleAndSaveToBuffer() {
	var output bytes.Buffer
	PrintConsoleAndSaveToBuffer(&output, "test")
	fmt.Println(output.String())
	// Output:
	// Printing without formatting: test
}

func TestPrintfConsoleAndSaveToBuffer(t *testing.T) {
	word := "test"
	want := `Printing with formatting:
	test`
	var output1 bytes.Buffer
	PrintfConsoleAndSaveToBuffer(&output1, word)
	if got := output1.String(); got != want {
		t.Errorf("PrintfConsoleAndSaveToBuffer(os.Stdout, test) = %q, want %q", got, want)
	}
}

func ExamplePrintfConsoleAndSaveToBuffer() {
	var output bytes.Buffer
	PrintfConsoleAndSaveToBuffer(&output, "test")
	fmt.Println(output.String())
	// Output:
	// Printing with formatting:
	//	test
}
