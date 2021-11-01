package bufferedstreams

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
Buffered streams:

A buffer is a region of space in the memory. It can be a fixed or a variable size buffer to read data from or write data to. The bytes built-in package provides Buffer structure type to construct a variable size buffer.
We can also use bytes.NewBuffer function which initializes a variable size buffer with initial buffer value. The bytes.NewBufferString function does the same thing but creates a new buffer with bytes of a string as its initial value.

	func NewBuffer(buf []byte) *bytes.Buffer
	func NewBufferString(s string) *bytes.Buffer

	https://pkg.go.dev/bytes#Buffer

*/

func BufferStreamsExample() {

	//create new buffer
	buf := bytes.NewBufferString("Hello world, ")

	// write some data to the buffer
	fmt.Print("bytes written => ")
	io.WriteString(buf, "How are you doing !\n")
	//buf.WriteTo(os.Stdout)
	buf.WriteString("Are you good !\n")

	p := make([]byte, 10)
	buf.Read(p)
	io.WriteString(os.Stdout, string(p))
}
