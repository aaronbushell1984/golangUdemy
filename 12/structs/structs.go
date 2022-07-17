// Package structs demonstrates the use of structs in go
package structs

import "fmt"

type food struct {
	colour string
}

type meat struct {
	food
	colour string
	fish   bool
}

// GetBananaFoodColour creates value of the type food and returns the field colour
//
// The struct food type is declared in the package:
//	type food struct {
//		colour string
//	}
// The value is declared of the type using:
//	banana := food{
//		colour: "yellow",
//	}
func GetBananaFoodColour() string {
	banana := food{
		colour: "yellow",
	}
	return banana.colour
}

// GetBananaType creates value of the type food and returns its type
func GetBananaType() string {
	banana := food{
		colour: "yellow",
	}
	return fmt.Sprintf("%T", banana)
}

// GetSalmonFoodType returns the type of the embedded type
// The struct meat type is declared in the package:
//	type meat struct {
//		food
//		fish bool
//	}
// The value is declared of the embedded type using:
//	salmon := meat{
//		food: food{
//			colour: "pink",
//		},
//		fish: true,
//	}
func GetSalmonFoodType() string {
	salmon := meat{
		food: food{
			colour: "pink",
		},
		fish: true,
	}
	return fmt.Sprintf("%T", salmon)
}

// GetSalmonFoodColour returns the field colour of embedded type
//
// The field colour can be accessed using dot notation either:
//	salmon.colour
// or:
//	salmon.food.colour
func GetSalmonFoodColour() string {
	salmon := meat{
		food: food{
			colour: "pink",
		},
		fish: true,
	}
	return fmt.Sprint(salmon.colour)
}

// GetFieldWhenCollision shows a name collision which should usually be avoided
//
// Declaring inner type:
//	type food struct {
//		colour string
//	}
// With outer type:
//	type meat struct {
//		food
//		colour string
//		fish   bool
//	}
// Causes a collision on colour:
// 	steak := meat{
//		food: food{
//			colour: "brown",
//		},
//		colour: "red",
//		fish:   false,
//	}
// Brown can be returned from inner type:
//	steak.food.colour
// Red can be returned from outer type:
//	steak.colour
// Avoid this by not declaring same field in outer type:
//	type meat struct {
//		food
//		fish   bool
//	}
// Colour will be promoted from inner type
func GetFieldWhenCollision(innerOrOuter string) string {
	steak := meat{
		food: food{
			colour: "brown",
		},
		colour: "red",
		fish:   false,
	}
	if innerOrOuter == "inner" {
		return steak.food.colour
	}
	return steak.colour
}

// GetAnonymousPineappleColour uses anonymous struct to declare and use
// Food type would be declared by:
//	type food struct {
//		colour string
//	}
// Pineapple could be assigned to food type using:
//	pineapple := food{
//		colour: "yellow",
//	}
// Anonymous struct does not use food type but declares the same by substituting exactly the same struct declaration where food could be:
//	pineapple := struct {
//		colour string
//	}{
//		colour: "yellow",
//	}
// This can help limit scope of struct if only required for one function
func GetAnonymousPineappleColour() string {
	pineapple := struct {
		colour string
	}{
		colour: "yellow",
	}
	return pineapple.colour
}
