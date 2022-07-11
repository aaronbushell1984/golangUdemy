// Package conditionals demonstrates the use of conditionals in go
package conditionals

// ReturnSuppliedNameIfOverThreeCharactersCase returns a string if it is over 3 characters in length
//
// Uses the case statement:
//	case len(name) < 4:
//		return "Name must be over three characters"
//	}
//	return name
func ReturnSuppliedNameIfOverThreeCharactersCase(name string) string {
	switch {
	case len(name) < 4:
		return "Name must be over three characters"
	}
	return name
}

// ReturnSuppliedNameIfOverThreeCharactersInitializationStatement returns a string if it is over 3 characters in length
//
// Uses an initialization statement:
//	if minimum := 4; len(name) < minimum {
//		return "Name must be over three characters"
//	}
//	return name
func ReturnSuppliedNameIfOverThreeCharactersInitializationStatement(name string) string {
	if minimum := 4; len(name) < minimum {
		return "Name must be over three characters"
	}
	return name
}

// FizzBuzzCaseFallthrough returns text based on supplied number:
//
// "Fizz Fallthrough" if divisible by 3
//
// "Buzz" if divisible by 5
//
// "FizzBuzz" if divisible by both 3 and 5
//
// "Nothing" if divisible by neither
//
// It demonstrates the use of case statements with a fallthrough:
//	switch {
//		case number%3 == 0 && number%5 == 0:
//			result = "FizzBuzz"
//		case number%3 == 0:
//			result = "Fizz"
//			fallthrough
//		case number == 2:
//			result += " Fallthrough"
//		case number%5 == 0:
//			result += "Buzz"
//		default:
//			result = "Nothing"
//	}
// The case for number == 2 is not evaluated when number%3 == 0 but falls through to " Fallthrough"
func FizzBuzzCaseFallthrough(number int) string {
	var result string
	switch {
	case number%3 == 0 && number%5 == 0:
		result = "FizzBuzz"
	case number%3 == 0:
		result = "Fizz"
		fallthrough
	case number == 2:
		result += " Fallthrough"
	case number%5 == 0:
		result += "Buzz"
	default:
		result = "Nothing"
	}
	return result
}
