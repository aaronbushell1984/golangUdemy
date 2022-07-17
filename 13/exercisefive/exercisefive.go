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
// Ranging over each person in turn appends their flavours to return list:
// 	for _, v := range mona.favouriteIceCreams {
//		returnIceCreams = append(returnIceCreams, v)
//	}
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
	for _, v := range ethan.favouriteIceCreams {
		returnIceCreams = append(returnIceCreams, v)
	}
	for _, v := range mona.favouriteIceCreams {
		returnIceCreams = append(returnIceCreams, v)
	}
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
// Using nested range loops to append to slice:
//	var iceCreams []string
//	for _, pers := range people {
//		for _, iceCream := range pers.favouriteIceCreams {
//			iceCreams = append(iceCreams, iceCream)
//		}
//	}
//	return iceCreams
func GetFavouriteIceCreamsFromMap(people map[string]person) []string {
	var iceCreams []string
	for _, pers := range people {
		for _, iceCream := range pers.favouriteIceCreams {
			iceCreams = append(iceCreams, iceCream)
		}
	}
	return iceCreams
}

// makeTruck constructs a new truck which is an outer type of vehicle
func makeTruck(drs int, clr string, fWDve bool) truck {
	return truck{
		vehicle{
			drs,
			clr,
		},
		fWDve,
	}
}

// makeSedan constructs a new sedan which is an outer type of vehicle
func makeSedan(drs int, clr string, lux bool) sedan {
	return sedan{
		vehicle{
			drs,
			clr,
		},
		lux,
	}
}
