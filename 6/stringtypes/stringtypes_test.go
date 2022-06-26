package stringtypes

import "testing"

func TestPrintStringAndTypeMrunalini(t *testing.T) {
	name := "Mrunalini"
	want := "Mrunalini is a string"
	if got := PrintStringAndType(name); got != want {
		t.Errorf("PrintStringAndType(name) = %q, want %q", got, want)
	}
}

func TestConvertStringToSliceOfBytesMrunalini(t *testing.T) {
	name := "Mrunalini"
	want := "[77 114 117 110 97 108 105 110 105]"
	if got := ConvertStringToSliceOfByteString(name); got != want {
		t.Errorf("ConvertStringToSliceOfBytes(name) = %q, want %q", got, want)
	}
}

func TestNewlineAndTabRawStringLiteral(t *testing.T) {
	want := "This is a:\n\t\t\t\traw string literal with new line and tabs"
	if got := NewlineAndTabRawStringLiteral(); got != want {
		t.Errorf("NewlineAndTabRawStringLiteral() = %q, want %q", got, want)
	}
}

func TestPrintStringCodepoints(t *testing.T) {
	name := "Mrunalini"
	want := "U+004D 'M' U+0072 'r' U+0075 'u' U+006E 'n' U+0061 'a' U+006C 'l' U+0069 'i' U+006E 'n' U+0069 'i' "
	if got := PrintStringCodepoints(name); got != want {
		t.Errorf("ConvertStringToSliceOfBytes(name) \ngot:\n %q, \nwant:\n %q", got, want)
	}
}
