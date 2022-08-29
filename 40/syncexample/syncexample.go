package syncexample

import "sync"

// Counter contains a state with value:
//
//	value int
//
// Channels are best used for passing data ownership, mutex is a great way of managing state:
//
//	mu    sync.Mutex
//
// NB. It is not advisable to embed a mutex as it allows public access to lock and unlock:
//
//	sync.Mutex
//
// NB. Counter contains a mutex so is best used with pointers to avoid copying. A constructor has therefore been included to encourage its use.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc locks the value, defers unlocking it and increments the value by one
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value/count
//
//	return c.value
func (c *Counter) Value() int {
	return c.value
}

// NewCounter returns the address of a newly created Counter
//
//	return &Counter{}
//
// NB. construction with declaration and assignment is not encouraged as a non pointer value may lead to passing a lock by value:
//
//	Please Avoid: counter := Counter{}
func NewCounter() *Counter {
	return &Counter{}
}
