// Package errorhandling demonstrates handling errors in go
package errorhandling

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// CheckFmtPrintError returns the error value if an error is thrown with fmt.Println
//
// Otherwise it returns the error with a preceeding message. The value also returns to console so Example receives this also despite it being thrown away with _:
//	_, err := fmt.Println("Value From Println")
func CheckFmtPrintError() string {
	_, err := fmt.Println("Value From Println")
	if err != nil {
		return fmt.Sprintln(err)
	}
	return fmt.Sprintf("There was no error so the err value is: %v", err)
}

// PrintlnCapture shows that Example tests capture a console output if any string is returned feom function
func PrintlnCapture() string {
	fmt.Println("I was sent to console but not returned and still captured in example")
	return "I was not sent to console but returned and also captured in example"
}

// ScanForAnswer allows user input to be placed at the provided address of type pointer string
//
// Errors are captured using:
//	_, err := fmt.Scan(answer)
//	if err != nil {
//		return fmt.Sprintf("There was an error: %v", err)
//	}
// The example demonstrates how this error would be received as the user input is not provided when example called
func ScanForAnswer(answer *string) string {
	_, err := fmt.Scan(answer)
	if err != nil {
		return fmt.Sprintf("There was an error: %v", err)
	}
	return *answer
}

// DisplayQuestion handles format for Question sent to user console
func DisplayQuestion(question string) string {
	fmt.Print("Question - " + question)
	return ""
}

// DisplayAnswer handles for the Answer displayed to the user
func DisplayAnswer(answer ...string) string {
	for _, v := range answer {
		fmt.Println("Your Answer - " + v)
	}
	return ""
}

// AskThreeQuestionsAndDisplayOutput used Scan, Answer and Question displays to capture and display some input
//
// The example demonstartes the result from receiving an error due to the test being run without user input possible
func AskThreeQuestionsAndDisplayOutput() {
	var ans1 string
	var ans2 string
	var ans3 string
	DisplayQuestion("Name: ")
	ans1 = ScanForAnswer(&ans1)
	DisplayQuestion("Fave Food: ")
	ans2 = ScanForAnswer(&ans2)
	DisplayQuestion("Fave Sport: ")
	ans3 = ScanForAnswer(&ans3)
	DisplayAnswer(ans1, ans2, ans3)
}

// Createfile makes a file with the provided name and contents
//
// Errors are collected in console. A file named noPermission.txt with read only permissions will return example fail
func CreateFile(name string, contents string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader(contents)
	io.Copy(f, r)
}

// ReadFile reads the content of the provided file into a string
//
// Calling for a missing file produces error to console as per example
func ReadFile(name string) string  {
	contents, err := os.ReadFile(name)
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(contents)
}

// OpenFileReadAndClose demonstrates safely opening, reading and deferring close of a file using ioutil
func OpenFileReadAndClose(name string) string {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Sprint(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(bs)
}
