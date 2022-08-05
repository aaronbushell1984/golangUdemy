// Package exercisesix is the sixth set of exercises from the course
package exercisesix

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last string
	age int
}

type vocal interface {
	Speak() (string, string, int)
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

// GetAnswerOfEverything returns the answer according to Hitchhikers guide to the galaxy
//
// Demonstrates function expression for excercise 1 in example:
//	foo := GetAnswerOfEverything()
//	fmt.Println(foo)
func GetAnswerOfEverything() int {
	return 42
}

// GetAnswerOfEverythingAndWords returns the answer as a number and words, according to Hitchhikers guide to the galaxy
//
// Demonstrates capturing two returns for excercise 1 in example:
// 	bar, bars := GetAnswerOfEverythingAndWords()
// 	fmt.Println(bar, bars)
func GetAnswerOfEverythingAndWords() (int, string) {
	return 42, "Fourty Two"
}

// GetSum returns the sum of any number of numbers; either individualy added like:
//	GetSum(1, 2, 3)
// Or by declaring a slice:
//	numbers := []int{1, 2, 3}
// and passing in a slice which is spread:
//	GetSum(numbers...)
func GetSum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

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

// Speak returns the first name , last name and age of person
func (prs person) Speak() (string, string, int) {
	return prs.first, prs.last, prs.age
}

// VocalSpeak returns the first name , last name and age of vocal type
func VocalSpeak(v vocal) (string, string, int) {
	return v.Speak()
}

// area returns area of a circle
func (cir circle) area() float64 {
	return cir.radius * cir.radius * math.Pi
}

// area returns srea of square
func (sq square) area() float64 {
	return sq.length * sq.length
}

// GetArea takes any shape and returns the area
func GetArea(s shape) float64 {
	return s.area()
}

// GetStringAnonymousFunc demonstrates an anonymous function:
//	func() {
//		anonymousResult = "I was created by anonymous function"
//	}() 
func GetStringAnonymousFunc() string {
	var anonymousResult string
	func() {
		 anonymousResult = "I was created by anonymous function"
	}() 
	return anonymousResult
}

// AssignFuncToVariable returns a string and is assigned to a variable in example
func AssignFuncToVariable() string {
	return "I was assigned to a variable and then called"
}

// ReturnFunc returns a function which returns a string:
// 	func ReturnFunc() func() string {
// 		return func() string {
// 			return "I was returned from one function, assigned to a variable and then called"
// 		}
// 	}
//
// The function is assigned to a variable and called in the example
func ReturnFunc() func() string {
	return func() string {
		return "I was returned from one function, assigned to a variable and then called"
	}
}

// UsedInCallback returns a function which returns a string can change the tense of the statement
//
// The function is used in the example to show it will be used in a callback
func UsedInCallback(tense string) func() string {
	return func() string {
		return fmt.Sprintf("I %v used in a callback", tense)
	}
}

// CallbackExample taked in a callback and allows the function to be modified as per example
func CallbackExample(callback func() string) string {
	return callback()
}

// Interate returns a function which counts a local variable which is closed inside the function:
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
// Each function expression can now keep a seperate count as per example
func Iterate() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

