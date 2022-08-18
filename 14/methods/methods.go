// Package methods demonstrates methods in go
package methods

type Food struct {
	name    string
	colour  string
	spoiled bool
}

// MakeFood contructs a variable of type Food
func MakeFood(nme string, clr string, spld bool) Food {
	return Food{
		nme,
		clr,
		spld,
	}
}

// Spoil attaches to Food allowing setting of spoiled field to true using:
//	Food.Spoil(banana)
func (f Food) Spoil() Food {
	f.spoiled = true
	return f
}
