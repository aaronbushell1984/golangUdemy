// Package exercisenine is the ninth set of exercises from the course
//
// A pointer receiver is set on the Person struct:
//	func (p *Person) Speak() string {
//		return p.name
//	}
// A human interface defines that any type which implements "Speak() string" is ALSO of type human:
//	type human interface {
//		Speak() string
//	}
// It is therefore possible to accept a person address:
//	SpeakName(&jb)
// into a function with a human argument:
//	func SpeakName(h human) string {
//		return h.Speak()
//	}
// It is NOT possible to pass in a person type as human has a pointer receiver:
//	SpeakName(jb) -- does not compile
package exercisenine

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex

type Person struct{
	name string
}

type human interface {
	Speak() string
}

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

// Speak attaches to Person as a pointer reciever and returns their name
func (p *Person) Speak() string {
	return p.name
}

// SpeakName accepts a human which allows pointer to a person but not person
func SpeakName(h human) string {
	return h.Speak()
}

// RaceConditionIncrementor creates a race condition as go routines are accessing the same memory space
func RaceConditionIncrementor() int {
	wg.Add(100)
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			res := counter
			runtime.Gosched()
			res++
			counter = res
			wg.Done()
		}()
		
	}
	wg.Wait()
	return counter
}

// RaceConditionIncrementorWithMutex solves race condition of RaceConditionIncrementor() using mutex
func RaceConditionIncrementorWithMutex() int {
	wg.Add(100)
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			res := counter
			res++
			counter = res
			mu.Unlock()
			wg.Done()
		}()		
	}
	wg.Wait()
	return counter
}

//  RaceConditionIncrementorWithAtomic solves race condition of RaceConditionIncrementor() using sync/atomic
func RaceConditionIncrementorWithAtomic() int64 {
	wg.Add(100)
	var counter int64
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		
	}
	wg.Wait()
	return counter
}

// GetOs returns the current operating system
func GetOs() string {
	return runtime.GOOS
}

// GetArch returns the current architecture
func GetArch() string {
	return runtime.GOARCH
}