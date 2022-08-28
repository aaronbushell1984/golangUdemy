package main

import (
	"github.com/aaronbushell1984/golangUdemy/36/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{
		Duration: mocking.ConfiguredSleep,
		Sleeps:   time.Sleep,
	}
	mocking.Countdown(os.Stdout, sleeper)
}
