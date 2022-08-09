// Package interfaces demonstrates the use of interfaces in gp
package interfaces

type Food struct {
	name    string
	colour  string
	spoiled bool
}

// perishable represents behaviour to spoil food
type Perishable interface {
	Spoil() Food
}

// MakeFood constructs a variable of type food
func MakeFood(nme string, clr string, spld bool) Food {
	return Food{
		nme,
		clr,
		spld,
	}
}

// Spoil attaches to food allowing setting of spoiled field to true using:
//	food.Spoil(banana)
func (f Food) Spoil() Food {
	f.spoiled = true
	return f
}
