package errorinfo

import "fmt"

func ExampleGetSquareRootWithCustomNewError() {
	res, err := GetSquareRootWithCustomNewError(16)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomNewError(0)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomNewError(-1)
	fmt.Println(res, err)
	// Output:
	// 4 <nil>
	// 0 zero or negative not allowed
	// 0 zero or negative not allowed
}

func ExampleGetSquareRootWithCustomPackageError() {
	res, err := GetSquareRootWithCustomPackageError(16)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomPackageError(0)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomPackageError(-1)
	fmt.Println(res, err)
	// Output:
	// 4 <nil>
	// 0 an error occured in the errorinfo pacakge
	// 0 an error occured in the errorinfo pacakge
}

func ExampleGetSquareRootWithCustomErrorf() {
	res, err := GetSquareRootWithCustomErrorf(16)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomErrorf(0)
	fmt.Println(res, err)
	res, err = GetSquareRootWithCustomErrorf(-1)
	fmt.Println(res, err)
	// Output:
	// 4 <nil>
	// 0 the provided number: 0 is a: float64 but is not allowed to less or equal to zero
	// 0 the provided number: -1 is a: float64 but is not allowed to less or equal to zero
}

func ExampleGetSquareRootWithStructError() {
	res, err := GetSquareRootWithStructError(16)
	fmt.Println(res, err)
	res, err = GetSquareRootWithStructError(0)
	fmt.Println(res, err)
	res, err = GetSquareRootWithStructError(-1)
	fmt.Println(res, err)
	// Output:
	// 4 <nil>
	// 0 struct error: Udemy, Aaron, not allowed to less or equal to zero
	// 0 struct error: Udemy, Aaron, not allowed to less or equal to zero
}

func ExampleGetTypeOfError() {
	fmt.Println(GetTypeOfError(errErrorinfo))
	// Output:
	// *errors.errorString
}





