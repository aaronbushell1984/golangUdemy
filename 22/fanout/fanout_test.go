package fanout

import (
	"fmt"
	"sort"
	"testing"
)


func ExamplePopulate() {
	c := make(chan int)
	var numbers []int
	go func() {
		Populate(c, 100)
	}()	
	for v := range c {
		numbers = append(numbers, v)
	}
	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleTimeConsumingWork() {
	fmt.Println(TimeConsumingWork(100))
	// Output:
	// 100
}

func BenchmarkTimeConsumingWork(b *testing.B) {
	TimeConsumingWork(100)
}

func ExampleFanOutIn() {
	var numbers []int
	numberPopulateChannel := make(chan int)
	receiveChannel := make(chan int)
	go Populate(numberPopulateChannel, 100)
	go FanOutIn(numberPopulateChannel, receiveChannel)
	for v := range receiveChannel {
		numbers = append(numbers , v)
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleConcurrentTimeConsumingWork() {
	fmt.Println(ConcurrentTimeConsumingWork(100))
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleSequentialTimeConsumingWork() {
	fmt.Println(SequentialTimeConsumingWork(100))
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func BenchmarkConcurrentTimeConsumingWork(b *testing.B) {
	ConcurrentTimeConsumingWork(100)
}

func BenchmarkSequentialTimeConsumingWork(b *testing.B) {
	SequentialTimeConsumingWork(100)
}

func ExampleThrottleFanOutIn() {
	var numbers []int
	numberPopulateChannel := make(chan int)
	receiveChannel := make(chan int)
	go Populate(numberPopulateChannel, 100)
	go ThrottleFanOutIn(numberPopulateChannel, receiveChannel, 10)
	for v := range receiveChannel {
		numbers = append(numbers , v)
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func ExampleThrottleConcurrentTimeConsumingWork() {
	fmt.Println(ThrottleConcurrentTimeConsumingWork(100, 10))
	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}

func BenchmarkThrottleConcurrentTimeConsumingWork(b *testing.B) {
	ThrottleConcurrentTimeConsumingWork(100, 10)
}

