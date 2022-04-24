package exampletypes

import (
	"testing"
)

func TestGetTypeReflectInt(t *testing.T) {
	number := 42
	want := "int"
	if got := GetTypeReflect(number); got != want {
		t.Errorf("GetTypeReflect() = %q, want %q", got, want)
	}
}

func TestGetTypeReflectString(t *testing.T) {
	word := "test"
	want := "string"
	if got := GetTypeReflect(word); got != want {
		t.Errorf("GetTypeReflectString() = %q, want %q", got, want)
	}
}

func TestGetTypePercentTInt(t *testing.T) {
	number := 42
	want := "int"
	if got := GetTypePercentT(number); got != want {
		t.Errorf("GetTypePercentT(number) = %q, want %q", got, want)
	}
}

func TestGetTypePercentTString(t *testing.T) {
	word := "test"
	want := "string"
	if got := GetTypePercentT(word); got != want {
		t.Errorf("GetTypePercentT(word) = %q, want %q", got, want)
	}
}