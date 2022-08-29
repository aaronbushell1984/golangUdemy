package syncexample

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 times leaves at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		asserCounter(t, counter, 3)
	})
	t.Run("runs safe concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		asserCounter(t, counter, want)
	})
}

func asserCounter(t *testing.T, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("got: %d want: %d", counter.Value(), want)
	}
}
