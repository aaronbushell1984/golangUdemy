package sort

import "fmt"

func ExampleSortIntsAscending() {
	numbers := []int{3, 2, 4, 1}
	fmt.Print(SortIntsAscending(numbers...))
	// Output:
	// [1 2 3 4]
}

func ExampleTestSortedInts() {
	unsorted := []int{3, 2, 4, 1}
	sorted := []int{1, 2, 3, 4}
	fmt.Println(TestSortedInts(unsorted...))
	fmt.Println(TestSortedInts(sorted...))
	// Output:
	// false
	// true
}

func ExampleStableSortPeopleAscendingByAgeName() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}
	fmt.Println(StableSortPeopleAscendingByAgeName(people))
	// Output:
	// [{Alice 25} {Bob 25} {Colin 25} {Elizabeth 25} {Alice 75} {Alice 75} {Bob 75} {Elizabeth 75}]
}

func ExampleSortPeopleAscendingByAgeName() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}
	fmt.Println(SortPeopleAscendingByAgeName(people))
	// Output:
	// [{Alice 25} {Elizabeth 25} {Bob 25} {Colin 25} {Alice 75} {Bob 75} {Elizabeth 75} {Alice 75}]
}

