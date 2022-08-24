// Package exerciseeleven is the eleventh set of exercises from the course
package exerciseeleven

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type ErrCustom struct {
	prompt string
}

type SqrtError struct {
	pack string
	meth string
	err  error
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

// ErrCustom_Error attaches to ErrCustom to a the package and the custom prompt
//
// This method allows a call of the method on any error type
func (e ErrCustom) Error() string {
	return fmt.Sprintf("exerciseeleven error: %v", e.prompt)
}

// GetEven returns even numbers and allows a custom error ie:
//	ceGe := ErrCustom{
//		prompt: "geteven error",
//	}
// to be passed in as type error (polymorphism):
//	evens, err := GetEven(numbers, ceGe)
// This is due to the errCustom struct having method signature:
//	Error() string
// which satisfies the standard library error interface and therefore implicitly implements the interface
func GetEven(numbers []int, e error) ([]int, string) {
	var evens []int
	for _, v := range numbers {
		if v%2 == 0 {
			evens = append(evens, v)
		}
		if v == 0 {
			return evens, e.Error()
		}
	}
	return evens, ""
}

// Assertion demonstrates that a type may be asserted from error to ErrCustom using:
//	e.(ErrCustom)
// This allows access to a field on the ErrCustom struct (prompt) via:
//	e.(ErrCustom).prompt
// Direct access of prompt from the error value would not work:
//	e.prompt
func Assertion(e error) string {
	return e.(ErrCustom).prompt
}

// SqrtError_Error returns a string with package, method and error attached to any struct value of type sqrtError
func (se SqrtError) Error() string {
	return fmt.Sprintf("-package-%v -method-%v -error-%v", se.pack, se.meth, se.err)
}

// Sqrt returns the square root unless the number is less than or equal to zero where a sqrtError is returned
func Sqrt(n float64) (float64, error) {
	if n <= 0 {
		e := errors.New("cannot supply a number less than or equal to zero")
		return 0, SqrtError{pack: "exerciseeleven", meth: "sqrt", err: e}
	}
	return math.Sqrt(n), nil
}
