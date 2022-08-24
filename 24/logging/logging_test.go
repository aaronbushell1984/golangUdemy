package logging

import "fmt"

func ExamplePrintNoFileError() {
	fmt.Println(PrintNoFileError("example.txt", "example.txt"))
	fmt.Println(PrintNoFileError("example.txt", "fail.txt"))
	// Output:
	// File Created Opened and Closed Successfully
	// Oh snap an error: open fail.txt: no such file or directory
}

func ExampleLogNoFileError() {
	fmt.Println(LogNoFileError("example.txt", "example.txt"))
	fmt.Println(LogNoFileError("example.txt", "fail.txt"))
	// Output:
	// File Created Opened and Closed Successfully
	// A file opening error was logged
}

func ExampleLogWithFileNoFileError() {
	fmt.Println(LogWithFileNoFileError("example.txt", "example.txt"))
	fmt.Println(LogWithFileNoFileError("log.txt", "fail.txt"))
	// Output:
	// File Created Opened and Closed Successfully
	// A file opening error was logged
}

func ExampleLogWithFileFatalNoFileError() {
	fmt.Println(LogWithFileFatalNoFileError("example.txt", "example.txt"))
	fmt.Println(LogWithFileFatalNoFileError("log.txt", "fail.txt"))
	// Output:
	// File Created Opened and Closed Successfully
	// log.Fatalln(some message, err) would exit program and prevent return
}

func ExampleLogWithFilePanicNoFileError() {
	fmt.Println(LogWithFilePanicNoFileError("example.txt", "example.txt"))
	fmt.Println(LogWithFilePanicNoFileError("log.txt", "fail.txt"))
	// Output:
	// File Created Opened and Closed Successfully
	// log.Panicln(some message, err) would panic program and prevent return
}

func ExamplePrintPanicNoFileError() {
	fmt.Println(PrintPanicNoFileError("example.txt", "example.txt"))
	// Output:
	// File Created Opened and Closed Successfully
}
