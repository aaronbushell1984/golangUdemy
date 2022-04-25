package numerictypes

import (
	"fmt"
	"testing"
)

func TestPrintDefaultTypeTabularDrive(t *testing.T) {
	var tests = []struct {
		a interface{}
		want string
	}{
		{4.2, "float64"},
		{3.1415926535897932384626433832795028, "float64"},
		{0, "int"},
		{-1, "int"},
		{-0.00000, "float64"},
	}

	for _, tt :=range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T){
			got := PrintDefaultType(tt.a)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}