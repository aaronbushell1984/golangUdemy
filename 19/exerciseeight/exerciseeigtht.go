// Package exerciseeight is the eighth set of exercises from the course
package exerciseeight

import (
	"encoding/json"
	"log"
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