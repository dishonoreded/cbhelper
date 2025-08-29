package lib

import (
	"encoding/base64"
)

func UnGzipWithBase64Decoding(src []byte) ([]byte, error) {
	decoded, err := Base64Decode(src)
	if err != nil {
		return nil, err
	}
	return UnGzip(decoded)
}

func Base64Decode(src []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(buf, src)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func Base64Encode(src []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)
	return buf, nil
}
