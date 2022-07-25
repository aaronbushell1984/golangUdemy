package polymorphism

import "fmt"

func ExampleBrewBeer() {
	g1 := grain{
		name: "marris",
	}
	g2 := grain{
		name: "naked oats",
	}
	h1 := hop{
		name: "citra",
		ibu:  4,
	}
	ipa := recipe{
		name: "ipa",
		grains: []grain{
			g1,
			g2,
		},
		hops: []hop{
			h1,
		},
		yeast: "kviek",
	}
	fmt.Println(BrewBeer(ipa))
	// Output:
	// 20

}
