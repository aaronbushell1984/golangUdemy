package fuzzexample

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	fmt.Print(Add(1, 2, 3))
	// Output:
	// 6
}

func FuzzAdd(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, ns int) {
		out := Add(ns)
		if
		t.Errorf("%q", out)
	})
}
