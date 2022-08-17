package exerciseeleven

import "fmt"

func ExampleMarshalPerson() {
	jb := person{
		First:		"James",
		Last:		"Bond",
		Sayings:	[]string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	fmt.Println(MarshalPerson(jb))
	// Output:
	// {"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]} <nil>
}

func ExampleJsonPerson() {
	jb := person{
		First:		"James",
		Last:		"Bond",
		Sayings:	[]string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	fmt.Println(JsonPerson(jb))
	// Output:
	// {"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]} <nil>
}

func ExampleGetEven() {
	numbers := []int{1, 2, 3, 4}
	evens, err := GetEven(numbers)
	fmt.Println(evens)
	// Output:
	// [2, 4]
}


