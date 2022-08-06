package polymorphism

import "fmt"

func ExampleMakeFood() {
	banana := MakeFood("banana", "yellow", false)
	fmt.Println(banana)
	// Output:
	// {banana yellow false}
}

func ExampleMakeFlower() {
	rose := MakeFlower("rose", "red", false)
	fmt.Println(rose)
	// Output:
	// {rose red false}
}

func ExampleSpoil() {
	banana := MakeFood("banana", "yellow", false)
	banana = food.Spoil(banana)
	rose := MakeFlower("rose", "red", false)
	rose = flower.Spoil(rose)
	fmt.Println(banana)
	fmt.Println(rose)
	// Output:
	// {banana yellow true}
	// {rose red true}
}

func ExampleTestIsFood() {
	apple := MakeFood("apple", "red", true)
	tulip := MakeFlower("tulip", "purple", false)
	ten := 10
	fmt.Println(TestIsFood(apple))
	fmt.Println(TestIsFood(tulip))
	fmt.Println(TestIsFood(ten))
	// Output:
	// true
	// false
	// false
}

func ExampleTestIsPerishable() {
	pineapple := MakeFood("pineapple", "yellow", true)
	daisy := MakeFlower("daisy", "white", false)
	ten := 10
	fmt.Println(TestIsPerishable(pineapple))
	fmt.Println(TestIsPerishable(daisy))
	fmt.Println(TestIsPerishable(ten))
	// Output:
	// true
	// false
	// false
}
