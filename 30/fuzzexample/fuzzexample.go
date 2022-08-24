// Package fuzzexample demonstrates use of a fuzz test in go
package fuzzexample

func Add(numbers ...int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}
