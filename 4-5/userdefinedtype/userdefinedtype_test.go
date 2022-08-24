package userdefinedtype

import "testing"

type mytypestring string

var myint mytypeint
var mystring mytypestring

func TestPrintType(t *testing.T) {
	want := "userdefinedtype.mytypeint"
	if got := PrintType(myint); got != want {
		t.Errorf("PrintType(myint) = %q, want %q", got, want)
	}
}

func TestPrintUnderlyingTypeInt(t *testing.T) {
	want := "int"
	if got := PrintUnderlyingType(myint); got != want {
		t.Errorf("PrintUnderlyingType(myint) = %q, want %q", got, want)
	}
}

func TestPrintUnderlyingTypeString(t *testing.T) {
	want := "string"
	if got := PrintUnderlyingType(mystring); got != want {
		t.Errorf("PrintUnderlyingType(myint) = %q, want %q", got, want)
	}
}

func TestPrintUserDefinedIntConversion(t *testing.T) {
	want := "int"
	if got := PrintUserDefinedIntConversion(myint); got != want {
		t.Errorf("PrintUserDefinedIntConversion(myint) = %q, want %q", got, want)
	}
}
