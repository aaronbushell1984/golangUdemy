package variables

import "testing"

func TestPrintReassignedVariables(t *testing.T) {
	want := 99
	if got := GetReassignedVariables(); got != want {
		t.Errorf("GetReassignedVariables() = %q, want %q", got, want)
	}
}

func TestPrintExpressionVariables(t *testing.T) {
	want := 124
	if got := GetExpressionVariables(); got != want {
		t.Errorf("GetExpressionVariables() = %q, want %q", got, want)
	}
}

func TestGetLocalVariables(t *testing.T) {
	want := "Available in function"
	if got := GetLocalVariables(); got != want {
		t.Errorf("GetLocalVariables() = %q, want %q", got, want)
	}
}

func TestGetPackageVariables(t *testing.T) {
	want := "Available outside function"
	if got := GetPackageVariables(); got != want {
		t.Errorf("GetPackageVariables() = %q, want %q", got, want)
	}
}

func TestGetPackageZeroValueInt(t *testing.T) {
	want := 0
	if got := GetPackageZeroValueInt(); got != want {
		t.Errorf("GetPackageZeroValueInt() = %q, want %q", got, want)
	}
}

func TestGetPackageZeroValueString(t *testing.T) {
	want := ""
	if got := GetPackageZeroValueString(); got != want {
		t.Errorf("GetPackageZeroValueInt() = %q, want %q", got, want)
	}
}
