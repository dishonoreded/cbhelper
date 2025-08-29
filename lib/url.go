package lib

import "net/url"

func URLEscape(raw []byte) ([]byte, error) {
	return []byte(url.QueryEscape(string(raw))), nil
}

func URLUnescape(raw []byte) ([]byte, error) {
	result, err := url.QueryUnescape(string(raw))
	return []byte(result), err
}
