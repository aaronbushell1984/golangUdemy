package exercisefive

import "fmt"

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
