package exerciseten

import (
	"fmt"
	"sort"
)

func ExampleReceiveFromChannelFuncLiteral() {
	fmt.Println(ReceiveFromChannelFuncLiteral())
	// Output:
	// 42
}

func ExampleReceiveFromBufferedChannel() {
	fmt.Println(ReceiveFromBufferedChannel())
	// Output:
	// 42
}

func ExampleGetValueInReceiveOnlyChannel() {
	fmt.Println(<-GetValueInReceiveOnlyChannel())
	// Output:
	// 42
}

func ExampleReceive() {
	c := Generate100()
	fmt.Println(Receive(c))
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleGenerate100() {
	numbers := Receive(Generate100())
	fmt.Println(numbers)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleReceiveWithSelect() {
	q := QuitWithZero()
	c := Generate100()
	numbers := ReceiveWithSelect(c, q)
	fmt.Println(numbers[len(numbers)-1])
	// Output:
	// 0
}

func ExampleQuitWithZero() {
	fmt.Println(<-QuitWithZero())
	// Output:
	// 0
}

func ExampleGetOkFromOpenChannel() {
	fmt.Println(GetOkFromOpenChannel())
	// Output:
	// true
}

func ExampleGetOkFromClosedChannel() {
	fmt.Println(GetOkFromClosedChannel())
	// Output:
	// The first ok is: true
	// The second ok is: false
}

func ExampleFanOut10GoRoutinesCountingToTen() {
	fmt.Println(<-FanOut10GoRoutinesCountingToTen())
	// Output:
	// 1
}

func ExampleFanIn() {
	out := FanOut10GoRoutinesCountingToTen()
	numbers := FanIn(out)
	sort.Ints(numbers)
	fmt.Println(numbers)
	// Output:
	// [1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 4 4 4 4 4 4 4 4 4 4 5 5 5 5 5 5 5 5 5 5 6 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 8 8 8 8 8 8 8 8 8 8 9 9 9 9 9 9 9 9 9 9 10 10 10 10 10 10 10 10 10 10]
}
