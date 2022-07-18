// Package methods demonstrates methods in go
package methods

type food struct {
	name    string
	colour  string
	spoiled bool
}

// MakeFood contructs a variable of type food
func MakeFood(nme string, clr string, spld bool) food {
	return food{
		nme,
		clr,
		spld,
	}
}

// FoodSpoil attaches to food allowing setting of spoiled field to true using:
//	food.FoodSpoil(banana)
func (f food) FoodSpoil() food {
	f.spoiled = true
	return f
}
