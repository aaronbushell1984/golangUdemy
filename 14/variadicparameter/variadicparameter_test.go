package variadicparameter

import "fmt"

func ExampleCombineSlices() {
	person1 := []string{"Mrunalini", "Bushell"}
	person2 := []string{"Ethan", "Bushell"}
	person3 := []string{"Luke", "Bushell"}
	numbers1 := []int{0, 1, 2, 3, 4, 5, 6}
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(CombineSlices(person1, person2, person3))
	fmt.Println(CombineSlices(numbers1, numbers2))
	// Output:
	// [[Mrunalini Bushell] [Ethan Bushell] [Luke Bushell]]
	// [[0 1 2 3 4 5 6] [1 2 3 4 5 6]]
}
