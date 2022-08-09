package interfaces

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

func ExampleFood_Spoil_persishable() {
	apple := MakeFood("apple", "red", true)
	apple = Perishable.Spoil(apple)
	fmt.Println(apple)
	// {apple red true}
}
