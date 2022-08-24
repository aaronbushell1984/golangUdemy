package tabletests

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{0, 2, 3}, 5},
		{[]int{-2, 2, 3}, 3},
	}
	for _, v := range tests {
		got := Sum(v.data...)
		if got != v.answer {
			t.Errorf("Got: %v Want: %v", got, v.answer)
		}
	}
}
