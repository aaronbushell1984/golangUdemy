package exercisefour

import "fmt"

func ExampleGetTypeOfArray() {
	arrInt := [5]int{0, 1, 2, 3, 4}
	arrString := [3]string{"Zero", "One", "Two"}
	arrFloat := [3]float64{0.0, 0.1, 0.2}
	fmt.Println(GetTypeOfArray(arrInt))
	fmt.Println(GetTypeOfArray(arrString))
	fmt.Println(GetTypeOfArray(arrFloat))
	// Output:
	// [5]int
	// [3]string
	// [3]float64
}

func ExampleGetTypeOfSlice() {
	arrInt := []int{0, 1, 2, 3, 4}
	arrString := []string{"Zero", "One", "Two"}
	arrFloat64 := []float64{0.0, 0.1, 0.2}
	arrFloat32 := []float32{0.0, 0.1, 0.2}
	fmt.Println(GetTypeOfSlice(arrInt))
	fmt.Println(GetTypeOfSlice(arrString))
	fmt.Println(GetTypeOfSlice(arrFloat64))
	fmt.Println(GetTypeOfSlice(arrFloat32))
	// Output:
	// []int
	// []string
	// []float64
	// []float32
}

func ExampleGetPartSlice() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(GetPartSlice(slice, 0, 4))
	fmt.Println(GetPartSlice(slice, 5, 9))
	fmt.Println(GetPartSlice(slice, 2, 6))
	fmt.Println(GetPartSlice(slice, 1, 5))
	// Output:
	// [42 43 44 45 46]
	// [47 48 49 50 51]
	// [44 45 46 47 48]
	// [43 44 45 46 47]
}

func ExampleAddToSlice() {
	intSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	newIntSlice := []int{52, 53, 54, 55, 56, 57, 58, 59, 60}
	fmt.Println(AddToSlice(intSlice, 52))
	fmt.Println(AddToSlice(intSlice, 52, 53, 54, 55))
	fmt.Println(AddToSlice(intSlice, newIntSlice...))
	// Output:
	// [42 43 44 45 46 47 48 49 50 51 52]
	// [42 43 44 45 46 47 48 49 50 51 52 53 54 55]
	// [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
}

func ExampleDeleteFromSliceByIndex() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(DeleteFromSliceByIndex(slice, 3, 6))
	// Output:
	// [42 43 44 48 49 50 51]
}

func ExampleGetSliceCapacity() {
	states := make([]string, 56, 57)
	fmt.Println(GetSliceCapacity(states))
	// Output:
	// 57
}

func ExampleGetSliceSize() {
	states := make([]string, 56, 57)
	fmt.Println(GetSliceSize(states))
	// Output:
	// 56
}

func ExampleGetIndexedMapFromSlice() {
	states := []string{"Alabama", "Alaska", "American Samoa", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District of Columbia", "Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Minor Outlying Islands", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Northern Mariana Islands", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Puerto Rico", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "U.S. Virgin Islands", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	fmt.Println(GetIndexedMapFromSlice(states, 56))
	// Output:
	// map[0:Alabama 1:Alaska 2:American Samoa 3:Arizona 4:Arkansas 5:California 6:Colorado 7:Connecticut 8:Delaware 9:District of Columbia 10:Florida 11:Georgia 12:Guam 13:Hawaii 14:Idaho 15:Illinois 16:Indiana 17:Iowa 18:Kansas 19:Kentucky 20:Louisiana 21:Maine 22:Maryland 23:Massachusetts 24:Michigan 25:Minnesota 26:Minor Outlying Islands 27:Mississippi 28:Missouri 29:Montana 30:Nebraska 31:Nevada 32:New Hampshire 33:New Jersey 34:New Mexico 35:New York 36:North Carolina 37:North Dakota 38:Northern Mariana Islands 39:Ohio 40:Oklahoma 41:Oregon 42:Pennsylvania 43:Puerto Rico 44:Rhode Island 45:South Carolina 46:South Dakota 47:Tennessee 48:Texas 49:U.S. Virgin Islands 50:Utah 51:Vermont 52:Virginia 53:Washington 54:West Virginia 55:Wisconsin 56:Wyoming]
}

func ExampleGetMultiDimensionalSlicesFromSlices() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hello, James"}
	numbers1 := []int{0, 1, 2, 3, 4, 5, 6}
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := []int{1, 2, 3, 4, 5}
	fmt.Println(GetMultiDimensionalSlicesFromSlices(jb, mp))
	fmt.Println(GetMultiDimensionalSlicesFromSlices(numbers1, numbers2, numbers3))
	// Output:
	// [[James Bond Shaken, not stirred] [Miss Moneypenny Hello, James]]
	// [[0 1 2 3 4 5 6] [1 2 3 4 5 6] [1 2 3 4 5]]
}

func ExampleAddRecordToMap() {
	quotes := make(map[string]string)
	quotes["Dude"] = "Will you just calm down"
	quotes["Walter"] = "Calmer than you are"
	character := "Donny"
	quote := "Phone's ringing Dude"
	fmt.Println(AddRecordToMap(quotes, character, quote))
	// Output:
	// map[Donny:Phone's ringing Dude Dude:Will you just calm down Walter:Calmer than you are]
}

func ExampleDataIsInMap() {
	map1 := map[string]string{
		"a": "A",
		"b": "B",
	}
	fmt.Println(DataIsInMap(map1, "a"))
	fmt.Println(DataIsInMap(map1, "c"))
	// Output:
	// true
	// false
}

func ExampleDeleteFromMapByKey() {
	quotes := make(map[string]string)
	quotes["Dude"] = "Will you just calm down"
	quotes["Walter"] = "Calmer than you are"
	DeleteFromMapByKey(quotes, "Walter")
	fmt.Println(quotes)
	// Output:
	// map[Dude:Will you just calm down]
}
