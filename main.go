package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var dataFlag string

func init() {
	flag.StringVar(&dataFlag, "msg", "hello", "the data you want to send")
	flag.Parse()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("error", err)
	}

	// I could simply use conn.Write , but there was a problem with NewWriter
	// I wasnt receiving because Writer doesnt not explicity Flush the data to conn.

	data := fmt.Sprintf("send %s from a go client", dataFlag)

	writer := bufio.NewWriter(conn)
	if _, err = writer.WriteString(data); err != nil {
		fmt.Println("Failed to write data, Err: ", err)
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Failed to flush the buffer, Err: ", err)
	}

	fmt.Println("data sent succefully")
}
