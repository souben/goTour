package writers

import (
	"io"
)

/*
io.Writer :
The io.Writer interface declares the basic Write method. Any type that
implements this interface is a type of Writer.

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

*/

type SampleStore struct {
	data []byte
}

func NewSampleStore(store []byte) *SampleStore {
	return &SampleStore{data: store}
}

// implement `io.Writer` interface
func (ss *SampleStore) Write(p []byte) (n int, err error) {

	if len(ss.data) == 10 {
		return 0, io.EOF
	}

	// calculate rest capacity in ss.data and copy it to a new slice
	remainingSize := 10 - len(ss.data)
	if remainingSize >= len(p) {
		n = len(p)
	} else {
		n = remainingSize
		err = io.EOF
	}
	ss.data = append(ss.data, p[:n]...)
	return
}

func (ss *SampleStore) GetDataFromStore() []byte {
	return ss.data
}
