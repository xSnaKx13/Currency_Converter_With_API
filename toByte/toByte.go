package tobyte

import "encoding/json"

func ToByte(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data, nil
}
