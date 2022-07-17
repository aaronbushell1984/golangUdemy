package structs

import "fmt"

func ExampleGetBananaFoodColour() {
	fmt.Println(GetBananaFoodColour())
	// Output:
	// yellow
}

func ExampleGetBananaType() {
	fmt.Println(GetBananaType())
	// Output:
	// structs.food
}

func ExampleGetSalmonFoodType() {
	fmt.Println(GetSalmonFoodType())
	// Output:
	// structs.meat
}

func ExampleGetSalmonFoodColour() {
	fmt.Println(GetSalmonFoodColour())
	// Output:
	// pink
}

func ExampleGetFieldWhenCollision() {
	fmt.Println(GetFieldWhenCollision("inner"))
	fmt.Println(GetFieldWhenCollision("outer"))
	// Output:
	// brown
	// red
}

func ExampleGetAnonymousPineappleColour() {
	fmt.Println(GetAnonymousPineappleColour())
	// Output:
	// yellow
}
