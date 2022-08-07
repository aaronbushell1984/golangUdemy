package exercisefive

import (
	"fmt"
)

func ExampleGetFavouriteIceCreams() {
	fmt.Println(GetFavouriteIceCreams())
	// Output:
	// [Strawberry Vanilla Chocolate Lemon]
}

func ExamplePlacePersonTypesInMap() {
	fmt.Println(PlacePersonTypesInMap())
	// Output:
	// map[Ethan:{Ethan Bushell [Strawberry Vanilla Chocolate]} Mrunalini:{Mrunalini Bushell [Lemon]}]
}

func ExampleGetAllMapPeopleInSlice() {
	people := PlacePersonTypesInMap()
	fmt.Println(GetAllMapPeopleInSlice(people))
	// Output:
	// [{Ethan Bushell [Strawberry Vanilla Chocolate]} {Mrunalini Bushell [Lemon]}]
}

func ExampleGetFavouriteIceCreamsFromMap() {
	people := PlacePersonTypesInMap()
	fmt.Println(GetFavouriteIceCreamsFromMap(people))
	// Output:
	// [Strawberry Vanilla Chocolate Lemon]
}

func ExamplemakeTruck() {
	monster := makeTruck(2, "red", true)
	fmt.Println(monster)
	fmt.Printf("%T\n", monster)
	fmt.Printf("The monster truck has 4WD: %v\n", monster.fourWheelDrive)
	fmt.Printf("The monster truck has %v doors\n", monster.doors)
	// Output:
	// {{2 red} true}
	// exercisefive.truck
	// The monster truck has 4WD: true
	// The monster truck has 2 doors
}

func ExamplemakeSedan() {
	cruiser := makeSedan(4, "black", true)
	fmt.Println(cruiser)
	fmt.Printf("%T\n", cruiser)
	fmt.Printf("The cruiser sedan is luxury: %v\n", cruiser.luxury)
	fmt.Printf("The cruiser sedan has %v doors\n", cruiser.doors)
	// Output:
	// {{4 black} true}
	// exercisefive.sedan
	// The cruiser sedan is luxury: true
	// The cruiser sedan has 4 doors
}

func ExamplemakeAnonymousSportsCar() {
	ferrari := makeAnonymousSportsCar(4, "red", false)
	fmt.Println(ferrari)
	fmt.Printf("%T\n", ferrari)
	// Output:
	// {{4 red} false}
	// struct { exercisefive.vehicle; nitrox bool }
}
