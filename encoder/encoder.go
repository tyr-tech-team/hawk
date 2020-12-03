package encoder

// Encoder -
type Encoder interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

//
var (
	JSON Encoder
	YAML Encoder
)

func init() {
	JSON = jsonEncoder{}
	YAML = yamlEncoder{}
}
