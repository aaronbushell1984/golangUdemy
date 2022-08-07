// Package exercisenine is the ninth set of exercises from the course
package exercisenine

import "sync"

var wg sync.WaitGroup

// TwoGoRoutines launches two go routines which append to a slice of strings:
//	go func() {
//		strings = append(strings, "I'm a go routine")
//		wg.Done()
//	}()
//	go func() {
//		strings = append(strings, "I'm a go routine")
//		wg.Done()
//	}()
// The waitgroup is set to 2:
//	wg.Add(2)
// Each routine reduces this waitgroup count by one when done with:
//	wg.Done()
// The function waits to return the strings until waitgroup is reduces twice:
//	wg.Wait()
//	return strings
func TwoGoRoutines() []string {
	var strings []string
	wg.Add(2)
	go func() {
		strings = append(strings, "I'm a go routine")
		wg.Done()
	}()
	go func() {
		strings = append(strings, "I'm a go routine")
		wg.Done()
	}()
	wg.Wait()
	return strings
}