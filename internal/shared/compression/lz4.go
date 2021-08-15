package compression

import (
	"bytes"
	Interfaces "diffme.dev/diffme-api/internal/core/interfaces"
	"encoding/gob"
	"github.com/pierrec/lz4"
)

type LZ4Compression struct {
	name string
}

func NewLZ4Compression() Interfaces.Compression {
	return &LZ4Compression{
		name: "LZ4",
	}
}

func (c *LZ4Compression) Compress(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)

	if err != nil {
		return nil, err
	}

	dataBytes := buf.Bytes()

	compressed := make([]byte, len(dataBytes))

	_, err = lz4.CompressBlock(dataBytes, compressed, nil)

	if err != nil {
		panic(err)
	}

	return compressed, nil
}

func (c *LZ4Compression) Decompress(compressed []byte) (interface{}, error) {
	decompressed := make([]byte, len(compressed))

	_, err := lz4.UncompressBlock(compressed, decompressed)

	if err != nil {
		panic(err)
	}

	return decompressed, nil
}
