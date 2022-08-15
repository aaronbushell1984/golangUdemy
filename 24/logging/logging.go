// Package logging demonstrates the use of logging and its difference to printing in go
package logging

import (
	"fmt"
	"log"
	"os"
)

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

func LogWithFileFatalNoFileError(create string, open string) string {
	f, err := os.Create(create)
	if err != nil {
		log.Fatalln("Oh snap an error: ", err)
		return "A file creation error was logged"
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open(open)
	if err != nil {
		log.Fatalln("Oh snap an error: ", err)
		return "A file opening error was logged"
	}
	defer f2.Close()
	return "File Created Opened and Closed Successfully"
}