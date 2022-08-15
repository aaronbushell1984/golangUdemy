package errorhandling

import (
	"fmt"
)

func ExampleCheckFmtPrintError() {
	fmt.Println(CheckFmtPrintError())
	// Output:
	// Value From Println
	// There was no error so the err value is: <nil>
}

func ExamplePrintlnCapture() {
	fmt.Println(PrintlnCapture())
	// Output:
	// I was sent to console but not returned and still captured in example
	// I was not sent to console but returned and also captured in example
}

func ExampleScanForAnswer() {
	var answer string
	fmt.Println(ScanForAnswer(&answer))
	// Output:
	// There was an error: EOF
}

func ExampleAskThreeQuestionsAndDisplayOutput() {
	AskThreeQuestionsAndDisplayOutput()
	// Output:
	// Question - Name: Question - Fave Food: Question - Fave Sport: Your Answer - There was an error: EOF
	// Your Answer - There was an error: EOF
	// Your Answer - There was an error: EOF
}

func ExampleDisplayQuestion() {
	question := "What is your name? "
	DisplayQuestion(question)
	// Output:
	// Question - What is your name?  
}

func ExampleDisplayAnswer() {
	answer := "James"
	fmt.Println(DisplayAnswer(answer))
	// Output:
	// Your Answer - James
}

func ExampleCreateFile() {
	CreateFile("example.txt", "Example Content")
	fmt.Println(ReadFile("example.txt"))
	CreateFile("noPermission.txt", "Example Content")
	// Output:
	// Example Content
	// open noPermission.txt: permission denied
}

func ExampleReadFile() {
	CreateFile("example.txt", "Example Content")
	fmt.Println(ReadFile("example.txt"))
	fmt.Println(ReadFile("missing.txt"))
	// Output:
	// Example Content
	// open missing.txt: no such file or directory
}

func ExampleOpenFileReadAndClose() {
	CreateFile("example.txt", "Example Content")
	fmt.Println(OpenFileReadAndClose("example.txt"))
	fmt.Println(OpenFileReadAndClose("missing.txt"))
	// Output:
	// Example Content
	// open missing.txt: no such file or directory
}











