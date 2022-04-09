package dependency

import "testing"

func TestHelloDependency(t *testing.T) {
    want := "Hello, world."
    if got := HelloDependency(); got != want {
        t.Errorf("HelloDependency() = %q, want %q", got, want)
    }
}