package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func UnGzip(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return resB.Bytes(), nil
}

func Gzip(data []byte) ([]byte, error) {
	fmt.Println(len(data))
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}

	fmt.Println(len(buf.Bytes()))

	return buf.Bytes(), nil
}
