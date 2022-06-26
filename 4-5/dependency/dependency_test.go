package dependency

import (
	"fmt"
	"testing"
)

func TestHelloDependency(t *testing.T) {
	want := "Hello, world."
	if got := HelloDependency(); got != want {
		t.Errorf("HelloDependency() = %q, want %q", got, want)
	}
}

func ExampleHelloDependency() {
	fmt.Println(HelloDependency())
	// Output:
	// Hello, world.
}
