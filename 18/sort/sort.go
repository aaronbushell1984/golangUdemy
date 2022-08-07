// Package sort demonstrates sorting in go
package sort

import "sort"

// SortIntsAscending sorts any number of numbers ascending into a slice using:
//	sort.Ints(numbers)
// This is the default behaviour
func SortIntsAscending(numbers ...int) []int {
	sort.Ints(numbers)
	return numbers
}

// TestSortedInts returns true if any number of ints are in ascending order using:
//	sort.IntsAreSorted(data)
func TestSortedInts(data ...int) bool {
	return sort.IntsAreSorted(data)
}

// StableSortPeopleAscendingByAge sorts people by name then age and guarantees stable sort
func StableSortPeopleAscendingByAgeName(people []struct{Name string; Age int}) []struct{Name string; Age int} {
	sort.SliceStable(people, func(i, j int) bool{
		return people[i].Name < people[j].Name
	})
	sort.SliceStable(people, func(i, j int) bool{
		return people[i].Age < people[j].Age
	})
	return people
}

// SortPeopleAscendingByAge sorts people by name then age and does not guarantee stable sort
func SortPeopleAscendingByAgeName(people []struct{Name string; Age int}) []struct{Name string; Age int} {
	sort.Slice(people, func(i, j int) bool{
		return people[i].Name < people[j].Name
	})
	sort.Slice(people, func(i, j int) bool{
		return people[i].Age < people[j].Age
	})
	return people
}
