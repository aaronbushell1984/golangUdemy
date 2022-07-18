package methods

import "fmt"

func ExampleMakeFood() {
	banana := MakeFood("banana", "yellow", false)
	fmt.Println(banana)
	// Output:
	// {banana yellow false}
}

func ExampleFoodSpoil() {
	banana := MakeFood("banana", "yellow", false)
	banana = food.FoodSpoil(banana)
	fmt.Println(banana)
	// Output:
	// {banana yellow true}
}
