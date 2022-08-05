package exercisesix

import "fmt"

func ExampleGetAnswerOfEverything() {
	foo := GetAnswerOfEverything()
	fmt.Println(foo)
	// Output:
	// 42
}

func ExampleGetAnswerOfEverythingAndWords() {
	bar, bars := GetAnswerOfEverythingAndWords()
	fmt.Println(bar, bars)
	// Output:
	// 42 Fourty Two
}

func ExampleGetSum() {
	fmt.Println(GetSum(1, 2, 3))
	numbers := []int{1, 2, 3}
	fmt.Println(GetSum(numbers...))
	// Output:
	// 6
	// 6
}

func ExampleGetFirstSecond() {
	fmt.Println(GetFirstSecond())
	// Output:
	// [first second]
}

func ExampleGetSecondFirst() {
	fmt.Println(GetSecondFirst())
	// Output:
	// [second]
}

func ExampleVocalSpeak() {
	james := person{
		first: "James",
		last: "Bond",
		age: 32,
	}
	fmt.Println(VocalSpeak(james))
	// Output:
	// James Bond 32
}

func ExampleGetArea() {
	circle1 := circle {
		radius: 4,
	}
	square1 := square {
		length: 4,
	}
	fmt.Println(GetArea(circle1))
	fmt.Println(GetArea(square1))
	// Output:
	// 50.26548245743669
	// 16
}

func ExampleGetStringAnonymousFunc() {
	fmt.Println(GetStringAnonymousFunc())
	// Output:
	// I was created by anonymous function
}

func ExampleAssignFuncToVariable() {
	assigned :=  AssignFuncToVariable()
	fmt.Println(assigned)
	// Output:
	// I was assigned to a variable and then called
}

func ExampleReturnFunc() {
	returned := ReturnFunc()
	fmt.Println(returned())
	// Output:
	// I was returned from one function, assigned to a variable and then called
}

func ExampleUsedInCallback() {
	fmt.Println(UsedInCallback("will be")())
	// Output:
	// I will be used in a callback
}


func ExampleCallbackExample() {
	fmt.Println(CallbackExample(UsedInCallback("was")))
	// Output:
	// I was used in a callback
}

func ExampleIterate() {
	first := Iterate()
	second := Iterate()
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(second())
	// Output:
	// 1
	// 2
	// 3
	// 1
}















