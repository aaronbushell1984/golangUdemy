package reflectexample

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice with 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("Sum() got: %v want: %v", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumAll() Got: %v Want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SumAllTails() Got: %v Want %v", got, want)
		}
	}
	t.Run("sum of tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
