// Package polymorphism demonstrates polymorphism in go
//
// Declaring the interface perishable defines the behaviour to spoil food:
//
//	type perishable interface {
//		Spoil() food
//	}
//
// Attaching the spoil method to type food:
//
//	func (f food) Spoil() food {
//		f.spoiled = true
//		return f
//	}
//
// Allows an apple to be of type food AND perishable:
//
//	func ExampleTestIsPerishable() {
//		apple := MakeFood("apple", "red", true)
//		ten := 10
//		fmt.Println(TestIsPerishable(apple))
//		fmt.Println(TestIsPerishable(ten))
//		// Output:
//		// true
//		// false
//		}
package polymorphism

type food struct {
	name    string
	colour  string
	spoiled bool
}

type flower struct {
	name    string
	colour  string
	spoiled bool
}

// perishable represents behaviour to spoil
type perishable interface {
	Spoil() (food, flower)
}

// MakeFood constructs a variable of type food
func MakeFood(nme string, clr string, spld bool) food {
	return food{
		name:    nme,
		colour:  clr,
		spoiled: spld,
	}
}

// MakeFlower constructs a variable of type flower
func MakeFlower(nme string, clr string, spld bool) flower {
	return flower{
		name:    nme,
		colour:  clr,
		spoiled: spld,
	}
}

// Spoil attaches to food allowing setting of spoiled field to true using:
//
//	food.Spoil(banana)
func (f food) Spoil() food {
	f.spoiled = true
	return f
}

// Spoil attaches to flower allowing setting of spoiled field to true using:
//
//	flower.Spoil(rose)
func (f flower) Spoil() flower {
	f.spoiled = true
	return f
}

// TestIsFood checks if the passed in variable is of type food using a type assertion:
//
//	t, ok := i.(T)
func TestIsFood(underTest interface{}) bool {
	_, ok := underTest.(food)
	return ok
}

// TestIsPerishable checks if the passed in variable is of type perishable using a type assertion:
//
//	t, ok := i.(T)
func TestIsPerishable(underTest interface{}) bool {
	_, ok := underTest.(perishable)
	return ok
}
