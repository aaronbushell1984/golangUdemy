// Package logging demonstrates the use of logging and its difference to printing in go
package logging

import (
	"fmt"
	"log"
	"os"
)

// PrintNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are returned via a print only:
//	return fmt.Sprintf("Oh snap an error: %v", err)
func PrintNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		return fmt.Sprintf("Oh snap an error: %v", err)
	}
	defer f.Close()
	f2, err := os.Open(open)
	if err != nil {
		return fmt.Sprintf("Oh snap an error: %v", err)
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}

// LogNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are logged to console and a string returned informing this occured:
//	log.Println("Oh snap an error: ", err)
//	return "A file opening error was logged"
func LogNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		log.Println("Oh snap an error: ", err)
		return "A file creation error was logged"
	}
	defer f.Close()
	f2, err := os.Open(open)
	if err != nil {
		log.Println("Oh snap an error: ", err)
		return "A file opening error was logged"
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}

// LogWithFileNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are logged to the first file:
//	log.SetOutput(f)
//	log.Println("Oh snap an error: ", err)
// and a string returned informing this occured:
//	return "A file opening error was logged"
func LogWithFileNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		log.Println("Oh snap an error: ", err)
		return "A file creation error was logged"
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open(open)
	if err != nil {
		log.Println("Oh snap an error: ", err)
		return "A file opening error was logged"
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}

// LogWithFileFatalNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are logged to the first file as fatal which would terminate the program if used:
//	log.SetOutput(f)
//	// log.Fatalln("Oh snap an error: ", err)
// and a string returned informing this would have occured:
//	return "log.Fatalln(some message, err) would exit program and prevent return"
func LogWithFileFatalNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		log.Fatalln("Oh snap an error: ", err)
		return "A file creation error was logged as fatal"
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open(open)
	if err != nil {
		// log.Fatalln("Oh snap an error: ", err)
		return "log.Fatalln(some message, err) would exit program and prevent return"
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}

// LogWithFilePanicNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are logged to the first file as panic which would panic the program if used:
//	log.SetOutput(f)
//	// log.Panicln("Oh snap an error: ", err)
// and a string returned informing this would have occured:
//	return "log.Panicln(some message, err) would panic program and prevent return"
func LogWithFilePanicNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		log.Panicln("Oh snap an error: ", err)
		return "A file creation error was logged as panic"
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open(open)
	if err != nil {
		// log.Panicln("Oh snap an error: ", err)
		return "log.Panicln(some message, err) would panic program and prevent return"
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}

// PrintPanicNoFileError creates a file named first argument and attempts to open file named the second argument
//
// Captured errors are printed in and program panics on any error:
//	log.SetOutput(f)
//	panic(err)
func PrintPanicNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		panic(err)

	}
	defer f.Close()
	f2, err := os.Open(open)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}
