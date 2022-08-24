// Package writerinterface demonstrates standard library use of the writer interface in go
//
// Several standard packages use the writer interface:
//	type Writer interface {
// 		Write(p []byte) (n int, err error)
//	}
// Any type which has the method:
//	Write(p []byte) (n int, err error)
// Is ALSO of type Writer
//
// Therefore both fmt.Println and io.WriteString can accept any writer
//
// This includes os.Stdout because it uses NewFile which has the same method with *file
package writerinterface

import (
	"fmt"
	"io"
)

// GetHelloIoWriterFprint ALSO captures the output which would be sent to console via os.Stdout
// 	fmt.Fprintln(buffer, "Hello")
func GetHelloIoWriterFprint(buffer io.Writer) {
	fmt.Fprintln(buffer, "Hello")
}

// GetHelloIoWriteString captures the output which would be sent to console via os.Stdout
//	io.WriteString(buffer, "Hello")
func GetHelloIoWriteString(buffer io.Writer) {
	io.WriteString(buffer, "Hello")
}
