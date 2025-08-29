package lib

import "encoding/json"

func FormatJSON(src []byte) []byte {
	var data map[string]interface{}
	err := json.Unmarshal(src, &data)
	if err != nil {
		panic(err)
	}

	result, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	return result
}
