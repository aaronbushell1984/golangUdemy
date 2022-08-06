// Package deferexample demonstartes the use of defer in go
package deferexample

// GetFirstSecond returns a string with first then second strings added
func GetFirstSecond() []string {
	var result []string
	first := func() {
		result = append(result, "first")
	}
	second := func() {
		result = append(result, "second")
	}
	first()
	second()
	return result
}

// GetSecondFirst is similar to GetFirstSecond() but because of defer:
//	defer first()
// the appended "first":
//	result = append(result, "first")
// is called after the return meaning "first" is missing in example
func GetSecondFirst() []string {
	var result []string
	first := func() {
		result = append(result, "first")
	}
	second := func() {
		result = append(result, "second")
	}
	defer first()
	second()
	return result
}
