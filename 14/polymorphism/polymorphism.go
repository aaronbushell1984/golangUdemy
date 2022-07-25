// Package polymorphism demonstrates polymorphism in go
package polymorphism

type hop struct {
	name string
	ibu  int
}

type grain struct {
	name string
}

type recipe struct {
	name   string
	grains []grain
	hops   []hop
	yeast  string
}

func BrewBeer(rec recipe) int {
	if rec.name == "ipa" {
		return 20
	}
	return 30
}
