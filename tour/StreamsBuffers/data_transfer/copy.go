package datatransfer

import (
	"io"
	"os"
	"strings"
)

/*
If we want to write data to an io.Writer object coming from
an io.Reader object, then we can use the io.Copy function.

	func Copy(dst Writer, src Reader) (int64, error)

	https://medium.com/rungo/introduction-to-streams-and-buffers-d148c0cda0ad

*/

func CopyExample() {
	// create a string io.Reader object
	stringReader := strings.NewReader("Holla hermano\n")

	// copy data from stringsReader to os.Stdout io.Writer
	io.CopyN(os.Stdout, stringReader, 4)

	buf := make([]byte, 1)
	io.CopyBuffer(os.Stdout, stringReader, buf)
}
