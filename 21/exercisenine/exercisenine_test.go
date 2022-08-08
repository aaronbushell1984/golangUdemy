package exercisenine

import "fmt"

func ExampleTwoGoRoutines() {
	fmt.Println(TwoGoRoutines())
	// Output:
	// [I'm a go routine I'm a go routine]
}

func ExamplePerson_Speak() {
	jb := Person{
		name: "James",
	}
	fmt.Println(jb.Speak())
	// Output:
	// James
}

func ExampleSpeakName() {
	jb := Person{
		name: "James",
	}
	fmt.Println(SpeakName(&jb))
	// Output:
	// James
}

func ExampleRaceConditionIncrementor() {
	fmt.Println(RaceConditionIncrementor() == 100)
	// Output:
	// false
}

func ExampleRaceConditionIncrementorWithMutex() {
	fmt.Println(RaceConditionIncrementorWithMutex() == 100)
	// Output:
	// true
}

func ExampleRaceConditionIncrementorWithAtomic() {
	fmt.Println(RaceConditionIncrementorWithAtomic() == 100)
	// Output:
	// true
}

func ExampleGetOs() {
	fmt.Println(GetOs())
	// Output:
	// linux
}

func ExampleGetArch() {
	fmt.Println(GetArch())
	// Output:
	// amd64
}
