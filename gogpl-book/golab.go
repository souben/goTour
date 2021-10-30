package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(conn)
		n, err := io.Copy(os.Stderr, conn)
		fmt.Println(conn)
		log.Printf("completed connection with %d bytes, err = %v", n, err)
	}
}
