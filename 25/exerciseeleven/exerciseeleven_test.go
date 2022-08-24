package exerciseeleven

import "fmt"

func ExampleMarshalPerson() {
	jb := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	fmt.Println(MarshalPerson(jb))
	// Output:
	// {"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]} <nil>
}

func ExampleJsonPerson() {
	jb := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	fmt.Println(JsonPerson(jb))
	// Output:
	// {"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]} <nil>
}

func ExampleErrCustom_Error() {
	ce := ErrCustom{
		prompt: "test custom error",
	}
	fmt.Println(ErrCustom.Error(ce))
	// Output:
	// exerciseeleven error: test custom error
}

func ExampleGetEven() {
	numbers := []int{1, 2, 3, 4}
	badNumbers := []int{0, 0, 0, 0}
	ceGe := ErrCustom{
		prompt: "geteven error",
	}
	badEvens, zeroErr := GetEven(badNumbers, ceGe)
	fmt.Println(badEvens, zeroErr)
	evens, err := GetEven(numbers, ceGe)
	fmt.Println(evens, err)
	// Output:
	// [0] exerciseeleven error: geteven error
	// [2 4]
}

func ExampleAssertion() {
	ceAe := ErrCustom{
		prompt: "asserted error",
	}
	fmt.Println(Assertion(ceAe))
	// Output:
	// asserted error
}

func ExampleSqrtError_Error() {
	se := SqrtError{
		pack: "exerciseelven_test",
		meth: "Example_SqrtError_Error",
		err:  nil,
	}
	fmt.Println(SqrtError.Error(se))
	// Output:
	// -package-exerciseelven_test -method-Example_SqrtError_Error -error-<nil>
}

func ExampleSqrt() {
	v1, err1 := Sqrt(16)
	v2, err2 := Sqrt(-10.23)
	fmt.Println(v1, err1)
	fmt.Println(v2, err2)
	// Output:
	// 4 <nil>
	// 0 -package-exerciseeleven -method-sqrt -error-cannot supply a number less than or equal to zero
}
