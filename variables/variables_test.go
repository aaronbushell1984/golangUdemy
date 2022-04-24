package variables

import "testing"

func TestPrintReassignedVariables(t *testing.T) {
	want := 99
	if got := PrintReassignedVariables(); got != want {
		t.Errorf("PrintReassignedVariables() = %q, want %q", got, want)
	}
}

func TestPrintExpressionVariables(t *testing.T) {
	want := 124
	if got := PrintExpressionVariables(); got != want {
		t.Errorf("PrintExpressionVariables() = %q, want %q", got, want)
	}
}