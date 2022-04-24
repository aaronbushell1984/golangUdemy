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

func TestGetGlobalVariables(t *testing.T) {
	want := "Available outside function"
	if got := GetGlobalVariables(); got != want {
		t.Errorf("GetGlobalVariables() = %q, want %q", got, want)
	}
}