package stringliteral

import "testing"

func TestGetStringLiteralWithQuotes(t *testing.T) {
	want := "When using quotes, quotation marks can not be printed and \nnewlines are interpreted with use of forward slash and n"
	if got := GetStringLiteralWithQuotes(); got != want {
		t.Errorf("GetStringLiteralWithBackticks() = %q, want %q", got, want)
	}
}