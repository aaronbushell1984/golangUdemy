// Package interfaces demonstrates the use of interfaces in gp
package interfaces

type food struct {
	name    string
	colour  string
	spoiled bool
}

// perishible represents behaviour to spoil food
type perishable interface {
	Spoil() food
}

// MakeFood contructs a variable of type food
func MakeFood(nme string, clr string, spld bool) food {
	return food{
		nme,
		clr,
		spld,
	}
}

// Spoil attaches to food allowing setting of spoiled field to true using:
//	food.Spoil(banana)
func (f food) Spoil() food {
	f.spoiled = true
	return f
}
