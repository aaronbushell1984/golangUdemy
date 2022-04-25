package main

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/4-5/capturereturns"
	"github.com/aaronbushell1984/golangUdemy/4-5/dependency"
	"github.com/aaronbushell1984/golangUdemy/4-5/exampleimport"
	"github.com/aaronbushell1984/golangUdemy/4-5/module"
	"github.com/aaronbushell1984/golangUdemy/4-5/variadicparameter"
)

func main() {
	fmt.Println("Main package and main function point is entry point in go")
	module.HelloModule()
	fmt.Println(dependency.HelloDependency(), "from the dependency module")
	exampleimport.ExampleImportPrint()
	variadicparameter.VaridaicParameterPrint()
	capturereturns.CaptureReturnsPrint()
}
