package main

import (
	"fmt"

	"github.com/aaronbushell1984/golangUdemy/21/exercisenine"
	"github.com/aaronbushell1984/golangUdemy/24/errorhandling"
	"github.com/aaronbushell1984/golangUdemy/4-5/dependency"
	"github.com/aaronbushell1984/golangUdemy/4-5/module"
	"github.com/aaronbushell1984/golangUdemy/4-5/printlnreturns"
	"github.com/aaronbushell1984/golangUdemy/6/stringtypes"
)

func main() {
	fmt.Println("Main package and main function point is entry point in go")
	module.HelloModule()
	fmt.Println(dependency.HelloDependency(), "from the dependency module")
	printlnreturns.PrintNumberAndError()
	s := "Mrunalini"
	fmt.Println(stringtypes.PrintStringCodepoints(s))
	fmt.Println(exercisenine.GetOs())
	fmt.Println(exercisenine.GetArch())
	errorhandling.AskThreeQuestionsAndDisplayOutput()
}
