// Package returnfunc demonstrates returning a function from a function
package returnfunc

func GetIntViaFunc() func() int {
	return func() int {
		return 42
	}
}