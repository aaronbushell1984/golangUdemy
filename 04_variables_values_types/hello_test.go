package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := HelloWorld(); got != want {
        t.Errorf("HelloWorld() = %q, want %q", got, want)
    }
}

func TestHelloDependency(t *testing.T) {
    want := "Hello, world."
    if got := HelloDependency(); got != want {
        t.Errorf("HelloDependency() = %q, want %q", got, want)
    }
}