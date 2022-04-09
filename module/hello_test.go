package module

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := HelloWorld(); got != want {
        t.Errorf("HelloWorld() = %q, want %q", got, want)
    }
}