package encoders

import (
	"bytes"
	"encoding/gob"
)

func EncodeJSON(data interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)

	if err != nil {
		return nil, err
	}

	dataBytes := buf.Bytes()

	return dataBytes, nil
}
