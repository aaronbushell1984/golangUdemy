package word

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/29/2/quote"
	"testing"
)

const q = quote.SunAlso

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(q)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(q)
	}
}

func ExampleCount() {
	fmt.Println(Count("One Two Three"))
	fmt.Println(Count(" One Two Three"))
	fmt.Println(Count(" One Two Three "))
	fmt.Println(Count(" One Two  Three "))
	fmt.Println(Count(q))
	// Output:
	// 3
	// 3
	// 3
	// 3
	// 1349
}

func ExampleUseCount() {
	s := " Written   Once Written Twice Written Thrice  "
	fmt.Println(UseCount(s))
	// Output:
	// map[Once:1 Thrice:1 Twice:1 Written:3]
}

func TestName(t *testing.T) {
	s := "    A   Tricky   Case  Don't you know...    "
	got := Count(s)
	want := 6
	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}

func TestUseCount(t *testing.T) {
	s := " Written   Once Written Twice Written Thrice  written again  "
	t.Run("three case sensitive words repeated returns 3", func(t *testing.T) {
		wordsMap := UseCount(s)
		gotWritten := wordsMap["Written"]
		assertCount(t, gotWritten, 3)
	})
	t.Run("lower case word present once", func(t *testing.T) {
		wordsMap := UseCount(s)
		got := wordsMap["written"]
		assertCount(t, got, 1)
	})
	t.Run("sum of totals correct", func(t *testing.T) {
		wordsMap := UseCount(q)
		var total int
		for _, v := range wordsMap {
			total += v
		}
		assertCount(t, total, 1349)
	})

}

func assertCount(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}
