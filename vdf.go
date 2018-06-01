package vdf

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/mdouchement/vdf/vdfg"
	"github.com/mdouchement/vdf/vdfp"
)

// Parse returns the unmarshalled Valve Data Format the given blob.
func Parse(blob []byte) (map[string]interface{}, error) {
	return vdfp.Parse(string(blob))
}

// ParseIO returns the unmarshalled Valve Data Format the given reader.
func ParseIO(r io.Reader) (map[string]interface{}, error) {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return vdfp.Parse(string(blob))
}

// Generate marshalizes the given data to Valve Data Format.
func Generate(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := vdfg.Generate(buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GenerateIO marshalizes the given data to Valve Data Format into w.
func GenerateIO(w io.Writer, data interface{}) error {
	return vdfg.Generate(w, data)
}
