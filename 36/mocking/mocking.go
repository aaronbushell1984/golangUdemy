// Package mocking demonstrates the use of spys, structs and interfaces to make mock tests
package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart  = 3
	finalWord       = "Go!"
	ConfiguredSleep = 5 * time.Second
)

// Sleeper defines behaviour to Sleep allowing a DefaultSleeper and ConfigurableSleeper
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is empty of fields but allows a method attachment
type DefaultSleeper struct{}

// Sleep attaches a Sleep of 1 second to receivers of DefaultSleeper
func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper exports Duration which is a time.duration and Sleeps which is function which accepts argument of time.duration
//
// Value of type ConfigurableSleeper implements sleep and therefore Sleeper interface
type ConfigurableSleeper struct {
	Duration time.Duration
	Sleeps   func(time.Duration)
}

// Sleep attaches to ConfigurableSleeper receiver passing the Duration into the Sleeps method
func (s *ConfigurableSleeper) Sleep() {
	s.Sleeps(s.Duration)
}

// Countdown writes to any io.Writer a countdown from configured countdownStart to 1 and then "Go!"
//
// Any value of type Sleeper must be passed which pause writing by appropriate duration
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}
