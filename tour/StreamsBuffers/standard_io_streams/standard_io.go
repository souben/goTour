package standardiostreams

import (
	"fmt"
	"io"
	"os"
)

/*
The standard I/O streams viz. os.Stdin, os.Stdout and os.Stderr
implement the io.Writer and io.StringWriter interfaces. Hence we
can call their Write or WriteString methods to write some data.
The fmt package also provides some functions to write some data
to io.Writer objects.

	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintln(w, a) (n int, err error)
	func Fprintf(w, format string, a) (n int, err error)

*/

func StandardIoExamples() {

	io.WriteString(os.Stdout, "Hello World!\n")
	os.Stdout.Write([]byte("Hello World"))

	// use `fmt` package function to write to a `io.Writer`
	fmt.Fprint(os.Stdout, "Hello World!\n")
	fmt.Fprintln(os.Stdout, "Hello World!") // adds new line
	fmt.Fprintf(os.Stdout, "%s, World!\n", "Hello")

}
