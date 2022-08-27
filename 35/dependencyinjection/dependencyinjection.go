// Package dependencyinjection demonstrates injecting io.Writer as a dependency to a function utilising:
//
//	fmt.Printf()
//
// As fmt.Printf uses:
//
//	fmt.Fprintf
//
// which in turn uses:
//
//	os.Stdout
//
// which implements interface:
//
//	io.Writer
//
// we can inject an io.Writer type into our function:
//
//	Greet(writer io.Writer, name string)
//
// and, therefore, add any value which implements writer interface into Greet
package dependencyinjection

import (
	fmt "fmt"
	"io"
)

// Greet signature started as:
//
//	func Greet(name string)
//
// To allow use of:
//
//	fmt.Fprintf(writer, "Hello, %s", name)
//
// inside the function, instead of:
//
//	fmt.Printf("Hello, %s", name)
//
// A writer of type *bytes.Buffer can be added as an argument:
//
//	func Greet(writer *bytes.Buffer, name string)
//
// As os.Stdout and bytes.Buffer both implement io.Writer was can instead use:
//
//	Greet(writer io.Writer, name string)
//
// We can now use Greet to pass messages to console by passing os.StdOut:
//
//	Greet(os.StdOut, "Chris")
//
// Or we could test this by passing in a buffer address
//
//	buffer := bytes.Buffer{}
//	Greet(&buffer, "Chris")
//	got := buffer.String()
func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return fmt.Errorf("error writing")
	}
	return nil
}
