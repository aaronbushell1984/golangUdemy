package fanin

import (
	"fmt"
	"sort"
	"sync"
)

func ExampleSendToOddAndEvenChannels() {
	odd := make(chan string)
	even := make(chan string)
	fanin := make(chan string)
	var numbers []string
	go SendToOddAndEvenChannels(even, odd)
	go FanInReceiveOddAndEvenChannel(even, odd, fanin)
	for v := range fanin {
		numbers = append(numbers, v)
	}
	sort.Strings(numbers)
	fmt.Println(numbers)
	// Output:
	// [Even: 10 Even: 2 Even: 4 Even: 6 Even: 8 Odd: 1 Odd: 3 Odd: 5 Odd: 7 Odd: 9]	
}

func ExampleFanInReceiveOddAndEvenChannel() {
	odd := make(chan string)
	even := make(chan string)
	fanin := make(chan string)
	var numbers []string
	go SendToOddAndEvenChannels(even, odd)
	go FanInReceiveOddAndEvenChannel(even, odd, fanin)
	for v := range fanin {
		numbers = append(numbers, v)
	}
	sort.Strings(numbers)
	fmt.Println(numbers)
	// Output:
	// [Even: 10 Even: 2 Even: 4 Even: 6 Even: 8 Odd: 1 Odd: 3 Odd: 5 Odd: 7 Odd: 9]
}

func ExampleSendBoringChannel() {
	var wg sync.WaitGroup
	wg.Add(1)
	in := SendBoringChannel("Bill")
	var strs []string
	go func() {
		for  v := range in {
			strs = append(strs, v)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(strs)
	// Output:
	// [0 Bill is boring 1 Bill is boring 2 Bill is boring]
}

func ExampleFanInToChannel() {
	var wg sync.WaitGroup
	wg.Add(1)
	ann := SendBoringChannel("Ann")
	bob := SendBoringChannel("Bob")
	fanIn := FanInToChannel(ann, bob)
	var strs []string
	go func() {
		for  v := range fanIn {
			strs = append(strs, v)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(strs)
	// Output:
	//
}



