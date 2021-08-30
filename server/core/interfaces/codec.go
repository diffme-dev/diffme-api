package interfaces

type Codec interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, p *interface{}) error
}
