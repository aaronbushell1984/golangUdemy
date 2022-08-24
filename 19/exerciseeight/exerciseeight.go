// Package exerciseeight is the eighth set of exercises from the course
package exerciseeight

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type agent struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// MarshalUser takes an array of user and returns the json in string format
func MarshalUser(users []user) string {
	userJson, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	return string(userJson)
}

// UnmarshallPeople takes the data and []person type to json convert to and return the person struct
func UnmarshallPeople(data []byte, p []person) []person {
	err := json.Unmarshal(data, &p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// EncodeJsonStdout encodes json directly to Stdout
//
// In encoding/json - NewEncoder accept the writer interface:
//	func NewEncoder(w io.Writer) *Encoder
// In this function - os.Stdout is also of Type Writer and therefore replaces (w io.Writer):
//	json.NewEncoder(os.Stdout)
// In encoding/json - NewEncoder returns a pointer to an Encoder:
//	func NewEncoder(w io.Writer) *Encoder
// In encoding/json - It can therefore be used as imput for Encode:
//	func (enc *Encoder) Encode(v any) error
// In this function - This can be and is called with dot notation to encode the agents:
//	json.NewEncoder(os.Stdout).Encode(agents)
func EncodeJsonStdout(agents []agent) {
	err := json.NewEncoder(os.Stdout).Encode(agents)
	if err != nil {
		log.Fatal(err)
	}
}

// SortNumbers sorts a slice of ints
func SortNumbers(numbers []int) {
	sort.Ints(numbers)
}

// SortStrings sorts a slice of strings
func SortStrings(strings []string) {
	sort.Strings(strings)
}

// SortPeopleAgeLast sorts people by age and then last name
func SortPeopleAgeLast(people []person) {
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Last < people[j].Last
	})
}

// SortPeopleSayings sorts the sayings of each person in people
func SortPeopleSayings(people []person) {
	for _, person := range people {
		sort.Strings(person.Sayings)
	}
}
