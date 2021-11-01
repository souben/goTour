package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	bufferedstreams "github.com/souben/streams_buffers/buffered_streams"
	datatransfer "github.com/souben/streams_buffers/data_transfer"
	"github.com/souben/streams_buffers/readers"
	standardiostreams "github.com/souben/streams_buffers/standard_io_streams"
	"github.com/souben/streams_buffers/writers"
)

func main() {

	fmt.Println(" ----- Stream Readers------")
	strData := readers.NewStringData("streams and buffers")
	p := make([]byte, 3)
	for {
		n, err := strData.Read(p)
		fmt.Printf("bytes : %d, data: %s\n", n, p[:n])
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			fmt.Println("end of streaming ...")
			break
		}
	}

	// using the built-in `strings.Newreader`
	src := strings.NewReader("streams and buffers")
	for {
		n, err := src.Read(p)
		fmt.Printf("bytes : %d, data: %s\n", n, p[:n])
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			fmt.Println("end of streaming ...")
			break
		}
	}

	fmt.Println(" ----- Stream Writers------")

	strStore := make([]byte, 0)
	simpleStore := writers.NewSampleStore(strStore)
	for {
		p := []byte("str")
		n, err := simpleStore.Write(p)
		if err == io.EOF {
			fmt.Printf("end of writing tto store : %s\n", simpleStore.GetDataFromStore())
			break
		} else if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Printf("wrote n bytes: %d, wrote data: %s, data in store: %s\n", n, p[:n], simpleStore.GetDataFromStore())
		}
	}

	anotherStore := writers.NewSampleStore(strStore)
	str := "hola !"
	n, err := io.WriteString(anotherStore, str)
	fmt.Printf("wrote n bytes: %d, data in 2 store: %s, err: %v\n", n, anotherStore.GetDataFromStore(), err)

	str = "hola mondo !"
	n, err = io.WriteString(anotherStore, str)
	fmt.Printf("wrote n bytes: %d, data in 2 store: %s, err: %v\n", n, anotherStore.GetDataFromStore(), err)

	str = "hola mondo !"
	n, err = io.WriteString(anotherStore, str)
	fmt.Printf("wrote n bytes: %d, data in 2 store: %s, err: %v\n", n, anotherStore.GetDataFromStore(), err)

	fmt.Println(" ----- Standard IO Streams ------")

	standardiostreams.StandardIoExamples()
	datatransfer.CopyExample()

	datatransfer.PipeExample()

	bufferedstreams.BufferStreamsExample()

	// create a buffered `io.Writer` from `Stdout`
	dst := bufio.NewWriter(os.Stdout)

	// write some data to the buffer
	fmt.Println(dst.WriteString("Hello World!"))
	fmt.Println(dst.Write([]byte(" How are you?\n")))

	// write buffered data to underlying `io.Writer`
	fmt.Println(dst.Flush(), "Flushed!")

	// write some more data to the buffer
	fmt.Println(dst.WriteString("Gooood?\n"))
	fmt.Println(dst.Flush(), "Flushed!")
}
