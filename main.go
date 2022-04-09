package main

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/dependency"
	"github.com/aaronbushell1984/golangUdemy/exampleimport"
	"github.com/aaronbushell1984/golangUdemy/module"
)

func main() {
	fmt.Println("Main package and main function point is entry point in go")
	module.HelloModule()
	fmt.Println(dependency.HelloDependency(), "from the dependency module")
	exampleimport.ExampleImportPrint()
}