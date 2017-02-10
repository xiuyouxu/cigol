package utils

import (
	"bytes"
	"io"
)

func Read2String(reader io.Reader) (string, error) {
	var buf bytes.Buffer
	buf.ReadFrom(reader)

	// efficient way to convert to string
	//	b := buf.Bytes()
	//	s := *(*string)(unsafe.Pointer(&b))

	ret := buf.String()

	return ret, nil
}
