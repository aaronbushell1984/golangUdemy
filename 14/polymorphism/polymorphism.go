// Package polymorphism demonstrates polymorphism in go
//
// Declaring the interface perishable defines the behaviour to spoil Food:
//
//	type perishable interface {
//		Spoil() Food
//	}
//
// Attaching the spoil method to type Food:
//
//	func (f Food) Spoil() Food {
//		f.spoiled = true
//		return f
//	}
//
// Allows an apple to be of type Food AND perishable:
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

type Food struct {
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
//
// Only one type can be placed in return:
//	Spoil() Food
// Therefore flower, even though it contains the spoil method is not perishable
type perishable interface {
	Spoil() Food
}

// MakeFood constructs a variable of type Food
func MakeFood(nme string, clr string, spld bool) Food {
	return Food{
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

// Spoil attaches to Food allowing setting of spoiled field to true using:
//
//	Food.Spoil(banana)
func (f Food) Spoil() Food {
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

// TestIsFood checks if the passed in variable is of type Food using a type assertion:
//	t, ok := i.(T)
func TestIsFood(underTest interface{}) bool {
	_, ok := underTest.(Food)
	return ok
}

// TestIsPerishable checks if the passed in variable is of type perishable using a type assertion:
//	t, ok := i.(T)
// NB. flower is NOT perishable because OR is not allowed in return type of interface:
//	type perishable interface {
//		Spoil() Food
//	}

func TestIsPerishable(underTest interface{}) bool {
	_, ok := underTest.(perishable)
	return ok
}
