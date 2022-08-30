package dog

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func ExampleYears() {
	fmt.Println(Years(0))
	fmt.Println(Years(4))
	fmt.Println(Years(-4))
	// Output:
	// 0 <nil>
	// 28 <nil>
	// 0 human years cannot be negative
}

func TestYears(t *testing.T) {
	_, err := Years(-4)
	want := errors.New("human years cannot be negative")
	if !reflect.DeepEqual(err, want) {
		t.Errorf("got: %q want: %q", err, want)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(0))
	fmt.Println(YearsTwo(4))
	fmt.Println(YearsTwo(-4))
	// Output:
	// 0 <nil>
	// 28 <nil>
	// 0 human years cannot be negative
}

func TestYearsTwo(t *testing.T) {
	_, err := YearsTwo(-4)
	want := errors.New("human years cannot be negative")
	if !reflect.DeepEqual(err, want) {
		t.Errorf("got: %q want: %q", err, want)
	}
}
