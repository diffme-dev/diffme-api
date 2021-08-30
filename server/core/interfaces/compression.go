package interfaces

type Compression interface {
	Decompress(data []byte) (interface{}, error)
	Compress(data interface{}) ([]byte, error)
}
