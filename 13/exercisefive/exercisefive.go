// Package exercisefive is the fifth set of exercises from the course
package exercisefive

type person struct {
	firstName          string
	lastName           string
	favouriteIceCreams []string
}

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheelDrive bool
}

type sedan struct {
	vehicle
	luxury bool
}

// GetFavouriteIceCreams combines and returns the favourite ice creams of two people
//
// Person is declared as a struct with a slice of ice creams:
//	type person struct {
//		firstName          string
//		lastName           string
//		favouriteIceCreams []string
//	}
// Two people are declared and assigned with ice creams:
//	ethan := person{
//		firstName: "Ethan",
//		lastName:  "Bushell",
//		favouriteIceCreams: []string{
//			"Strawberry",
//			"Vanilla",
//			"Chocolate",
//		},
//	}
//	mona := person{
//		firstName: "Mrunalini",
//		lastName:  "Bushell",
//		favouriteIceCreams: []string{
//			"Lemon",
//		},
//	}
// Appending to ice cream slice each persons favourites which are spread with ...:
//	returnIceCreams = append(returnIceCreams, mona.favouriteIceCreams...)
func GetFavouriteIceCreams() []string {
	ethan := person{
		firstName: "Ethan",
		lastName:  "Bushell",
		favouriteIceCreams: []string{
			"Strawberry",
			"Vanilla",
			"Chocolate",
		},
	}
	mona := person{
		firstName: "Mrunalini",
		lastName:  "Bushell",
		favouriteIceCreams: []string{
			"Lemon",
		},
	}
	var returnIceCreams []string
	returnIceCreams = append(returnIceCreams, ethan.favouriteIceCreams...)
	returnIceCreams = append(returnIceCreams, mona.favouriteIceCreams...)
	return returnIceCreams
}

// PlacePersonTypesInMap declares and assigns two persons and places them in a map called people
//
// Placing two person types in a people map:
//	people := map[string]person{
//		ethan.firstName: ethan,
//		mona.firstName:  mona,
//	}
func PlacePersonTypesInMap() map[string]person {
	ethan := person{
		firstName: "Ethan",
		lastName:  "Bushell",
		favouriteIceCreams: []string{
			"Strawberry",
			"Vanilla",
			"Chocolate",
		},
	}
	mona := person{
		firstName: "Mrunalini",
		lastName:  "Bushell",
		favouriteIceCreams: []string{
			"Lemon",
		},
	}
	people := map[string]person{
		ethan.firstName: ethan,
		mona.firstName:  mona,
	}
	return people
}

// GetAllMapPeopleInSlice demonstrates taking a map and appending to a slice:
//	var slicePeople []person
//	for _, v := range people {
//		slicePeople = append(slicePeople, v)
//	}
func GetAllMapPeopleInSlice(people map[string]person) any {
	var slicePeople []person
	for _, v := range people {
		slicePeople = append(slicePeople, v)
	}
	return slicePeople
}

// GetFavouriteIceCreamsFromMap returns a slice all ice creams from given people map
//
// Using range loop to append to slice...:
//	var iceCreams []string
//	for _, pers := range people {
//		iceCreams = append(iceCreams, pers.favouriteIceCreams...)
//	}
//	return iceCreams
func GetFavouriteIceCreamsFromMap(people map[string]person) []string {
	var iceCreams []string
	for _, pers := range people {
		iceCreams = append(iceCreams, pers.favouriteIceCreams...)
	}
	return iceCreams
}

// MakeTruck constructs a new truck which is an outer type of vehicle
func MakeTruck(drs int, clr string, fWDve bool) truck {
	return truck{
		vehicle{
			drs,
			clr,
		},
		fWDve,
	}
}

// MakeSedan constructs a new sedan which is an outer type of vehicle
func MakeSedan(drs int, clr string, lux bool) sedan {
	return sedan{
		vehicle{
			drs,
			clr,
		},
		lux,
	}
}

// MakeAnonymousSportsCar constructs a car with a new field anonymously, for example:
//	ferrari := makeAnonymousSportsCar(4, "red", false)
// Its fields can not be accessed outside the function like:
// 	ferrari.doors
func MakeAnonymousSportsCar(drs int, clr string, nos bool) any {
	any := struct {
		vehicle
		nitrox bool
	}{
		vehicle{
			drs,
			clr,
		},
		nos,
	}
	return any
}
