package main

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/29/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	humanEq, err := dog.Years(-10)
	if err != nil {
		fmt.Errorf("error running dog.Years(): %q", err)
	}
	fido := canine{
		name: "Fido",
		age:  humanEq,
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
