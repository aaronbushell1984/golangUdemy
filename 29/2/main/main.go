package main

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/29/2/quote"
	"github.com/aaronbushell1984/golangUdemy/29/2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
