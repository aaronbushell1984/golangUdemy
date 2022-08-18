package methods

import "fmt"

func ExampleMakeFood() {
	banana := MakeFood("banana", "yellow", false)
	fmt.Println(banana)
	// Output:
	// {banana yellow false}
}

func ExampleFood_Spoil() {
	banana := MakeFood("banana", "yellow", false)
	banana = Food.Spoil(banana)
	fmt.Println(banana)
	// Output:
	// {banana yellow true}
}
