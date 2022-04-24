package exampletypes

import (
	"testing"
)

func TestGetTypeReflectInt(t *testing.T) {
	number := 42
	want := "int"
	if got := GetTypeReflect(number); got != want {
		t.Errorf("GetTypeReflect() = %q, want %q", got, want)
	}
}

func TestGetTypeReflectString(t *testing.T) {
	word := "test"
	want := "string"
	if got := GetTypeReflect(word); got != want {
		t.Errorf("GetTypeReflectString() = %q, want %q", got, want)
	}
}

func TestGetTypePercentTInt(t *testing.T) {
	number := 42
	want := "int"
	if got := GetTypePercentT(number); got != want {
		t.Errorf("GetTypePercentT(number) = %q, want %q", got, want)
	}
}

func TestGetTypePercentTString(t *testing.T) {
	word := "test"
	want := "string"
	if got := GetTypePercentT(word); got != want {
		t.Errorf("GetTypePercentT(word) = %q, want %q", got, want)
	}
}

func TestGetRawStringWithBackticks(t *testing.T) {
	want := `When using backticks "quotes" are printed and
	newlines are interpreted without use of \n`
	if got := GetRawStringWithBackticks(); got != want {
		t.Errorf("GetRawStringWithBackticks() = %q, want %q", got, want)
	}
}

func TestGetStringLiteralWithQuotes(t *testing.T) {
	want := "When using quotes, quotation marks can not be printed and \nnewlines are interpreted with use of forward slash and n"
	if got := GetStringLiteralWithQuotes(); got != want {
		t.Errorf("GetStringLiteralWithBackticks() = %q, want %q", got, want)
	}
}