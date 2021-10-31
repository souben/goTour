package readers

import (
	"io"
)

// a strcut to to hold string data
type stringData struct {
	str       string
	readIndex int // default: 0
}

func NewStringData(str string) *stringData {
	return &stringData{str: str}
}

// add Read method (to implmenet io.Reader)
func (strData *stringData) Read(p []byte) (n int, err error) {

	// convert `str` to a slices of bytes
	byteData := []byte(strData.str)

	// if `read` is GTE source data length, return io.EOF error
	if strData.readIndex >= len(byteData) {
		return 0, io.EOF
	}

	// get nextRead Index
	nextReadIndex := strData.readIndex + len(p)

	if nextReadIndex >= len(byteData) {
		nextReadIndex = len(byteData)
		err = io.EOF
	}

	//get next bytes to copy and set n to its length
	nextBytes := byteData[strData.readIndex:nextReadIndex]
	n = len(nextBytes)

	//copy all bytes of nextBytes to p
	copy(p, nextBytes)

	// increment readIndex to nextReadLimit
	strData.readIndex = nextReadIndex

	// return values
	return
}

/*
ioutil.ReadAll :   func ReadAll(src io.Reader) ([]byte, error)

If you want to read everything from a source which implements io.Reader interface,
then you can use ioutil.ReadAll method.



io.ReadFull : func ReadFull(src Reader, buf []byte) (n int, err error)

If you are interested in exactly x number of bytes from the source and somehow,
if x number of bytes could not be read, get an error instead, io.ReadFull is
the function for you.


io.LimitReader: func LimitReader(r Reader, n int64) Reader

Many at times, we need to set a cap on a data source. The io.LimitReader
takes a Reader object r and returns a Reader object with a cap of n bytes.

*/
