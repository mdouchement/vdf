package vdfg

import (
	"errors"
	"io"
)

// IndentChar is the character that is used to indent the Valve Data Format.
const IndentChar = "\t"

// ErrUnsupportedType is return when the generator does not handle a datatype.
var ErrUnsupportedType = errors.New("Unsupported type")

// Generate marshalizes the given data to Valve Data Format.
func Generate(w io.Writer, data interface{}) error {
	return marshal(data, w, 0)
}

func marshal(data interface{}, w io.Writer, indent int) (err error) {
	switch x := data.(type) {
	case string:
		addindent(w, indent) // TODO handle writes errors

		io.WriteString(w, "\"")
		io.WriteString(w, x)
		io.WriteString(w, "\"")
	case map[string]interface{}:
		for k, v := range x {
			marshal(k, w, indent)

			switch y := v.(type) {
			case string:
				io.WriteString(w, "\t\t")
				marshal(y, w, 0)
			case map[string]interface{}:
				io.WriteString(w, "\n")
				addindent(w, indent)
				io.WriteString(w, "{\n")
				marshal(y, w, indent+1)
				io.WriteString(w, "\n")
				addindent(w, indent)
				io.WriteString(w, "}")
			default:
				return ErrUnsupportedType
			}
			io.WriteString(w, "\n")
		}
	default:
		return ErrUnsupportedType
	}

	return nil
}

func addindent(w io.Writer, indent int) error {
	for i := 0; i < indent; i++ {
		_, err := io.WriteString(w, IndentChar)
		if err != nil {
			return err
		}
	}
	return nil
}
