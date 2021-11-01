package datatransfer

import (
	"fmt"
	"io"
)

func PipeExample() {

	//create a pipe
	src, dst := io.Pipe()

	//start a goroutine that writes data to dst
	go func() {
		dst.Write([]byte("DATA_1")) //WRITE AND BLOCK
		dst.Write([]byte("ERTY_2"))
		dst.Close()
	}()

	// data transfer packet
	packet := make([]byte, 4)

	//read from src
	bytesRead1, err1 := src.Read(packet)
	fmt.Printf("BytesRead: %d, packet: %s, error: %v\n", bytesRead1, packet, err1)

	//read from src
	bytesRead2, err2 := src.Read(packet)
	fmt.Printf("BytesRead: %d, packet: %s, error: %v\n", bytesRead2, packet, err2)

	// read from src
	//read from src
	bytesRead3, err3 := src.Read(packet)
	fmt.Printf("BytesRead: %d, packet: %s, error: %v\n", bytesRead3, packet, err3)

	//read from src
	bytesRead3, err3 = src.Read(packet)
	fmt.Printf("BytesRead: %d, packet: %s, error: %v\n", bytesRead3, packet, err3)

	//read from src
	bytesRead3, err3 = src.Read(packet)
	fmt.Printf("BytesRead: %d, packet: %s, error: %v\n", bytesRead3, packet, err3)

}
