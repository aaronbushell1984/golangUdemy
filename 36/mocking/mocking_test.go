package mocking

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	Calls []string
}

func (o *SpyCountdownOperations) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

func (o *SpyCountdownOperations) Write(_ []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("output correct", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := fmt.Sprintf(expectedOutput())
		if got != want {
			t.Errorf("got: %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperWriter := &SpyCountdownOperations{}
		Countdown(spySleeperWriter, spySleeperWriter)
		want := expectedWriteSleepCalls()
		if !reflect.DeepEqual(want, spySleeperWriter.Calls) {
			t.Errorf("got: %q want %q", spySleeperWriter.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := ConfiguredSleep
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept: %v but actually slept: %v", sleepTime, spyTime.durationSlept)
	}
}

func expectedWriteSleepCalls() []string {
	var want []string
	for i := 0; i < countdownStart; i++ {
		want = append(want, write)
		want = append(want, sleep)
	}
	want = append(want, write)
	return want
}

func expectedOutput() string {
	var output string
	for i := countdownStart; i > 0; i-- {
		output += fmt.Sprintf("%v\n", i)
	}
	output += finalWord
	return output
}
