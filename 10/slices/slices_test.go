package slices

import (
	"fmt"
	"testing"
)

func ExampleReturnSlice() {
	strings := []string{"Slices", "are", "composite", "literals", "and hold values of same type"}
	integers := []int{1, 2, 3}
	notSlice := 1
	fmt.Println(ReturnSlice(strings))
	fmt.Println(ReturnSlice(integers))
	fmt.Println(ReturnSlice(notSlice))
	// Output:
	// [Slices are composite literals and hold values of same type]
	// [1 2 3]
	// This is not a slice
}

func ExampleGetFromSliceByIndex() {
	integers := []int{1, 2, 3}
	fmt.Println(GetFromSliceByIndex(integers, 2))
	fmt.Println(GetFromSliceByIndex(integers, 0))
	// Output:
	// 3
	// 1
}

func ExampleGetAllFromSlice() {
	integers := []int{1, 2, 3}
	moreIntegers := []int{-11, 2, 333, 9}
	fmt.Println(GetAllFromSlice(integers))
	fmt.Println(GetAllFromSlice(moreIntegers))
	// Output:
	// [1 2 3]
	// [-11 2 333 9]
}

func ExampleGetSliceOfSlice() {
	integersEasy := []int{0, 1, 2, 3, 4, 5}
	integersCaution := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GetSliceOfSlice(integersEasy, 0, 4))
	fmt.Println(GetSliceOfSlice(integersCaution, 0, 4))
	// Output:
	// [0 1 2 3]
	// [1 2 3 4]
}

func ExampleGetAppendedSlices() {
	slice1 := []int{0, 1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GetAppendedSlices(slice1, slice2))
	fmt.Println(GetAppendedSlices(slice2, slice1))
	// Output:
	// [0 1 2 3 4 5 1 2 3 4 5 6]
	// [1 2 3 4 5 6 0 1 2 3 4 5]
}

func ExampleGetDeletedSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(GetDeletedSlice(slice, 0, 2, 5, 6))
	// Output:
	// [0 1 2 5 6]
}

func ExampleGetCapacityOfSlice() {
	madeSlice := make([]int, 10, 100)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(GetCapacityOfSlice(madeSlice))
	fmt.Println(GetCapacityOfSlice(slice))
	// Output:
	// 100
	// 7
}

func ExampleGetLengthOfSlice() {
	madeSlice := make([]int, 10, 100)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(GetLengthOfSlice(madeSlice))
	fmt.Println(GetLengthOfSlice(slice))
	// Output:
	// 10
	// 7
}

func BenchmarkGetAppendedSlice(b *testing.B) {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	GetAppendedSlice(slice, 6, 9, 9, 1000)
}

func BenchmarkGetAppendedMadeSlice(b *testing.B) {
	madeSlice := make([]int, 7, 11)
	GetAppendedSlice(madeSlice, 6, 9, 9, 1000)
}

func ExampleGetAppendedSlice() {
	madeSlice := make([]int, 4, 4)
	fmt.Println(GetAppendedSlice(madeSlice, 6, 9, 9, 1000))
	// Output:
	// [0 0 0 0 6 9 9 1000]
}

func ExampleGetCombinedMultiDimensionalSlices() {
	person1 := []string{"Mrunalini", "Bushell"}
	person2 := []string{"Ethan", "Bushell"}
	person3 := []string{"Luke", "Bushell"}
	numbers1 := []int{0, 1, 2, 3, 4, 5, 6}
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GetCombinedMultiDimensionalSlices(person1, person2, person3))
	fmt.Println(GetCombinedMultiDimensionalSlices(numbers1, numbers2))
	// Output:
	// [[Mrunalini Bushell] [Ethan Bushell] [Luke Bushell]]
	// [[0 1 2 3 4 5 6] [1 2 3 4 5 6]]
}

func ExampleDataIsInMap() {
	map1 := map[int]string{
		65: "A",
		66: "B",
	}
	fmt.Println(DataIsInMap(map1, 65))
	fmt.Println(DataIsInMap(map1, 91))
	// Output:
	// true
	// false
}

func ExampleGetAsciiGlyphByDecimal() {
	fmt.Println(GetAsciiGlyphByDecimal(65))
	// Output:
	// A
}

func ExampleGetMapOfAsciiGlyphsByDecimal() {
	var uppers []int
	for i := 65; i <= 90; i++ {
		uppers = append(uppers, i)
	}
	fmt.Println(GetMapOfAsciiGlyphsByDecimal(65, 66))
	fmt.Println(GetMapOfAsciiGlyphsByDecimal(uppers...))
	// Output:
	// map[65:A 66:B]
	// map[65:A 66:B 67:C 68:D 69:E 70:F 71:G 72:H 73:I 74:J 75:K 76:L 77:M 78:N 79:O 80:P 81:Q 82:R 83:S 84:T 85:U 86:V 87:W 88:X 89:Y 90:Z]
}

func ExampleDeleteFromMapByKey() {
	letters := map[int]string{
		65: "A",
		66: "B",
		67: "C",
	}
	DeleteFromMapByKey(letters, 66)
	fmt.Println(letters)
	DeleteFromMapByKey(letters, 66)
	fmt.Println(letters)
	// Output:
	// map[65:A 67:C]
	// map[65:A 67:C]
}
