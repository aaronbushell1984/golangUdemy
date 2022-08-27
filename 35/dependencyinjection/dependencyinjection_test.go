package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	err := Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	if err != nil {
		t.Fatalf("error writing")
	}
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
