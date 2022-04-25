package exercises

import "testing"

type mytypeint int
var myint mytypeint
var regularint int

func TestGetShortDeclarations(t *testing.T) {
	want := "42 James Bond true"
	if got := GetShortDeclarations(); got != want {
		t.Errorf("GetShortDeclarations() = %q, want %q", got, want)
	}
}

func TestGetVarDeclarationZeroValues(t *testing.T) {
	want := `0 false`
	if got := GetVarDeclarationZeroValues(); got != want {
		t.Errorf("GetVarDeclarationZeroValues() = %q, want %q", got, want)
	}
}

func TestGetVarDeclarationsAsString(t *testing.T) {
	want := "42 James Bond true"
	if got := GetVarDeclarationsAsString(); got != want {
		t.Errorf("GetVarDeclarationsAsString() = %q, want %q", got, want)
	}
}

func TestPrintValueOfUserDefinedTypeInt(t *testing.T) {
	want := "0"
	if got := PrintValueOfUserDefinedType(myint); got != want {
		t.Errorf("PrintValueOfUserDefinedType() = %q, want %q", got, want)
	}
}

func TestPrintTypeOfUserDefinedTypeInt(t *testing.T) {
	want := "exercises.mytypeint"
	if got := PrintTypeOfUserDefinedType(myint); got != want {
		t.Errorf("PrintTypeOfUserDefinedType() = %q, want %q", got, want)
	}
}

func TestPrintValueOfUserDefinedTypeIntSetTo42(t *testing.T) {
	myint = 42
	want := "42"
	if got := PrintValueOfUserDefinedType(myint); got != want {
		t.Errorf("PrintValueOfUserDefinedType(myint) = %q, want %q", got, want)
	}
}

func TestPrintTypeOfUserDefinedConvertedToUnderlyingTypeInt(t *testing.T) {
	regularint = int(myint)
	want := "int"
	if got := PrintTypeOfUserDefinedType(regularint); got != want {
		t.Errorf("(PrintTypeOfUserDefinedConvertedToUnderlyingTypeInt(converted) = %q, want %q", got, want)
	}
}
