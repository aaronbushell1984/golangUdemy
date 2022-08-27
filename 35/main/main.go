// Package main demonstrates use of Greet function from package dependencyinjection
//
// As:
//
//	http.ResponseWriter
//
// also implements interface:
//
//	io.Writer
//
// then the same function can be used:
//
//	dependencyinjection.Greet(w, "world")
//
// which in turn can listen and serve "Hello, World" at localhost:5001
//
//	http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler
package main

import (
	"github.com/aaronbushell1984/golangUdemy/35/dependencyinjection"
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	dependencyinjection.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
