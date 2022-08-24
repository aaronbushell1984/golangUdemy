package operators

import (
	"fmt"
	"testing"
)

// func TestCheckEqualsFalse7And42(t *testing.T) {
// 	want := false
// 	if got := CheckEquals(7, 42); got != want {
// 		t.Errorf("CheckEquals(7, 42) = %v, want %v", got, want)
// 	}
// }

// func TestCheckEqualsTrue7And7(t *testing.T) {
// 	want := true
// 	if got := CheckEquals(7, 7); got != want {
// 		t.Errorf("CheckEquals(7, 7) = %v, want %v", got, want)
// 	}
// }

// refactor to table driven test
func TestCheckEqualsTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want bool
	}{
		{7, 42, false},
		{7, 7, true},
		{-1, 1, false},
		{0, 0, true},
		{7000000, 7000000, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := CheckEquals(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckGreaterThanTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want bool
	}{
		{7, 42, false},
		{7, 7, false},
		{1, -1, true},
		{0, 0, false},
		{7000002, 7000001, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := CheckGreaterThan(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
