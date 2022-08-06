package writerinterface

import (
	"bytes"
	"fmt"
)

func ExampleGetHelloIoWriterFprint() {
	var output bytes.Buffer
	GetHelloIoWriterFprint(&output)
	fmt.Println(output.String())
	// Output:
	// Hello
}

func ExampleGetHelloIoWriteString() {
	var output1 bytes.Buffer
	GetHelloIoWriteString(&output1)
	fmt.Println(output1.String())
	// Output:
	// Hello
}
