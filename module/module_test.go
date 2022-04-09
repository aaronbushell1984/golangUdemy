package module

import "testing"

func TestHelloModule(t *testing.T) {
    want := "Hello from module package"
    if got := HelloModule(); got != want {
        t.Errorf("HelloModule() = %q, want %q", got, want)
    }
}