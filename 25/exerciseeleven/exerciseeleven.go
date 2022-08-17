// Package exerciseeleven is the eleventh set of exercises from the course
package exerciseeleven

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First	string
	Last	string
	Sayings	[]string
}

type errCustom struct {
	prompt 	string
	err		error
}

// MarshalPerson takes a value of type person struct and encodes to a string
//
// It demonstrates handling the potential error:
//	if err != nil {
//		return fmt.Sprintf("There was an error: %v", err)
//	}
func MarshalPerson(p person) (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "Failed", fmt.Errorf("there was an error: %v", err)
	}
	return string(bs), err
}

// JsonPerson marshals a person to json and demonstrates using Errof from fmt:
//	return "Failed", fmt.Errorf("there was an error: %v", err)
func JsonPerson(p person) (string, error) {
	bs, err := MarshalPerson(p)
	if err != nil {
		return "Failed", fmt.Errorf("there was an error: %v", err)
	}
	return string(bs), err
}

func (e errCustom) GetEven(numbers []int) ([]int, error) {
	Error()
}
