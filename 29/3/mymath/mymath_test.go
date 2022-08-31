package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	avgTests := []struct {
		name    string
		numbers []int
		want    float64
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 3},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -3},
		{"mixed order", []int{5, 2, 1, 3, 4}, 3},
		{"mixed positive/negative", []int{-5, 2, -1, 4}, 0.5},
		{"with zero", []int{1, 2, 3, 4, 5, 0}, 2.5},
	}
	for _, tt := range avgTests {
		t.Run(tt.name, func(t *testing.T) {
			got := CenteredAvg(tt.numbers)
			if got != tt.want {
				t.Errorf("got: %f want: %f", got, tt.want)
			}
		})
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(numbers)
	}
}

func ExampleCenteredAvg() {
	numbers := []int{-5000, 1, 2, 3, 4, 5, 968}
	fmt.Println(CenteredAvg(numbers))
	// Output:
	// 3
}
