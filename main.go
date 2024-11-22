package main

import (
	"errors"
	"io"
)

func compareSlices(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func Contains(r io.Reader, seq []byte) (bool, error) {
	buff := make([]byte, 39)
	n, err := r.Read(buff)
	if err != nil {
		return false, errors.New("SUIIIIIIII")
	}
	if len(seq) > n {
		return false, errors.New("file is smaller then seq")
	}
	for i := 0; i < n-len(seq); i++ {
		if compareSlices(seq, buff[i:i+len(seq)]) {
			return true, nil
		}
	}
	return false, errors.New("sui")
}
