// Pacakge marshal demonstrates marshalling json in go
package marshal

import (
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

// MakePerson makes a Person by passing the fields it requires
func MakePerson(f string, l string, a int) Person {
	made := Person{
		First: f,
		Last: l,
		Age: a,
	}
	return made
}

// AddPerson adds any number of the type Person to the people slice
func AddPerson(ppl []Person, pp ...Person) []Person {
	ppl = append(ppl, pp...)
	return ppl
}

// Marshal uses json.Marshal to convert Person to json as string and error
func MarshalAsString(p []Person) (string, error) {
	bs, err := json.Marshal(p)
	return string(bs), err
}

// Marshal uses json.Marshal to convert Person to json as []byte and error
func MarshalAsByte(p []Person) ([]byte, error) {
	bs, err := json.Marshal(p)
	return bs, err
}