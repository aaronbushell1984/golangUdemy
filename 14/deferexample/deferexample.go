// Package deferexample demonstartes the use of defer in go
package deferexample

import "fmt"

// GetFirstSecond returns a string with first then second strings added
func GetFirstSecond() []string {
	var result []string
	first := func() {
		result = append(result, "first")
	}
	second := func() {
		result = append(result, "second")
	}
	first()
	second()
	return result
}

// GetSecondFirst is similar to GetFirstSecond() but because of defer:
//	defer first()
// the appended "first":
//	result = append(result, "first")
// is called after the return meaning "first" is missing in example
func GetSecondFirst() []string {
	var result []string
	first := func() {
		result = append(result, "first")
	}
	second := func() {
		result = append(result, "second")
	}
	defer first()
	second()
	return result
}

// FirstInLastOut shows defer executes statements in reverse order
func FirstInLastOut() string {
	var numbers string
	for i := 1; i < 6; i++ {
		defer fmt.Println(i)
	}
	return numbers
}

// DeferCapturedAtTimeOfCall shows that defer evaluates at time of call:
//	i := 0
//	defer fmt.Println(i)
// Even when variable i is incremented following this evaluation it does not change the previous evaluation and gives i as 0
func DeferCapturedAtTimeOfCall() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}
