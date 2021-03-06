package methods

import "fmt"

func ExampleMakeFood() {
	banana := MakeFood("banana", "yellow", false)
	fmt.Println(banana)
	// Output:
	// {banana yellow false}
}

func ExampleSpoil() {
	banana := MakeFood("banana", "yellow", false)
	banana = food.Spoil(banana)
	fmt.Println(banana)
	// Output:
	// {banana yellow true}
}
