package pointers

import "fmt"

func ExampleGetAddressOfValue() {
	value := "Value"
	fmt.Printf("%T", GetAddressOfValue(value))
	// Output:
	// *interface {}
}

func ExampleAssignToAddress() {
	ban := "banana"
	fmt.Printf("%T\n", ban)
	fmt.Printf("%T", AssignToAddress(ban))
	// Output:
	// string
	// string
}

func ExampleGetDereference() {
	apple := "apple"
	var addressApple *string = &apple
	fmt.Println(GetDereference(addressApple))
	// Output:
	// apple
}

func ExampleChangeValueViaNewIdentifier() {
	oldIdentifier := "old"
	fmt.Println(oldIdentifier)
	fmt.Println(ChangeValueViaNewIdentifier(oldIdentifier))
	// Output:
	// old
	// new
}

func ExampleGetChanged() {
	Same := "Same"
	fmt.Println(Same)
	fmt.Println(GetChanged(Same))
	fmt.Println(Same)
	// Output:
	// Same
	// Changed
	// Same
}

func ExampleGetPointerChanged() {
	Same := "Same"
	fmt.Println(Same)
	fmt.Println(GetPointerChanged(&Same))
	fmt.Println(Same)
	// Output:
	// Same
	// Changed
	// Changed
}
