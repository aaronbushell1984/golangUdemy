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











