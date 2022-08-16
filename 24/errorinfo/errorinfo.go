// Package errorinfo demonstrates techniques for adding information to errors in go
//
// A package wide error can be added, for example:
//	var errErrorinfo = errors.New("an error occured in the errorinfo pacakge")
// The convention is to name this with err preceeding package or function
//	errErrorinfo
// A custom error struct can be created, for example:
//	type errErrorinfoStruct struct {
//		course string
//		student string
//		err error
//	}
// when the Error() method is attached to the type:
//	
package errorinfo

import (
	"errors"
	"fmt"
	"math"
)

var errErrorinfo = errors.New("an error occured in the errorinfo pacakge")

type errErrorinfoStruct struct {
	course string
	student string
	err error
}

func (er errErrorinfoStruct) Error() string {
	return fmt.Sprintf("struct error: %v, %v, %v", er.course, er.student, er.err)
}

// GetSquareRootWithCustomNewError creates a new error if number is less or equal to zero
//	if number <= 0 {
//		return 0, errors.New("zero or negative not allowed")
//	}
func GetSquareRootWithCustomNewError(number float64) (float64, error) {
	if number <= 0 {
		return 0, errors.New("zero or negative not allowed")
	}
	return math.Sqrt(number), nil
}

// GetSquareRootWithCustomPackageError creates a package errErrorinfo error if number is less or equal to zero
func GetSquareRootWithCustomPackageError(number float64) (float64, error) {
	if number <= 0 {
		return 0, errErrorinfo
	}
	return math.Sqrt(number), nil
}

// GetSquareRootWithCustomErrorf uses format printing Errorf to return an error
func GetSquareRootWithCustomErrorf(number float64) (float64, error) {
	if number <= 0 {
		return 0, fmt.Errorf("the provided number: %v is a: %T but is not allowed to less or equal to zero", number, number)
	}
	return math.Sqrt(number), nil
}

// GetSquareRootWithStructError uses a custom struct with attached error method to return an error
//
// The error struct implicitly implements the standard library error interface because of the error method. This, therefore, also shows polymorphism
func GetSquareRootWithStructError(number float64) (float64, error) {
	if number <= 0 {
		return 0, errErrorinfoStruct{"Udemy", "Aaron", fmt.Errorf("not allowed to less or equal to zero")}
	}
	return math.Sqrt(number), nil
}

// GetTypeOfError returns the type of an error
func GetTypeOfError(er error) string {
	return fmt.Sprintf("%T", er)
}
