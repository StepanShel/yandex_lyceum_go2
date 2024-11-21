package main

import (
	"errors"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	buff := make([]byte, 39)
	n, err := r.Read(buff)
	if err != nil {
		return false, errors.New("SUIIIIIIIIIIIII")
	}
	if len(seq) < n {
		return false, errors.New("file is smaller then seq")
	}
	//for i := 0; i < n; i++{
	//if == buff[i:i+3]{

	//}
	//}
	return true, nil
}
func main() {
}
