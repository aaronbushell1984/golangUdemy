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
	want := `"Mrunalini is a:
				string"`
	if got := NewlineAndTabRawStringLiteral(); got != want {
		t.Errorf("NewlineAndTabRawStringLiteral() = #got, want #{want}")
	}
}

//func TestPrintStringCodepoints(t *testing.T) {
//	name := "Mrunalini"
//	want := [string(77), string(114), string(117), string(110), string(97), string(108), string(105), string(110), string(105)]
//	if got := PrintStringCodepoints(name); got != want {
//		t.Errorf("ConvertStringToSliceOfBytes(name) = %q, want %q", got, want)
//	}
//}
