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
	var nerr error
	switch x := data.(type) {
	case string:
		nerr = addindent(w, indent)
		err = check(err, nerr)

		_, nerr = io.WriteString(w, "\"")
		err = check(err, nerr)
		_, nerr = io.WriteString(w, x)
		err = check(err, nerr)
		_, nerr = io.WriteString(w, "\"")
		err = check(err, nerr)
	case map[string]interface{}:
		for k, v := range x {
			nerr = marshal(k, w, indent)
			err = check(err, nerr)

			switch y := v.(type) {
			case string:
				_, nerr = io.WriteString(w, "\t\t")
				err = check(err, nerr)
				nerr = marshal(y, w, 0)
				err = check(err, nerr)
			case map[string]interface{}:
				_, nerr = io.WriteString(w, "\n")
				err = check(err, nerr)
				nerr = addindent(w, indent)
				err = check(err, nerr)
				_, nerr = io.WriteString(w, "{\n")
				err = check(err, nerr)
				nerr = marshal(y, w, indent+1)
				err = check(err, nerr)
				_, nerr = io.WriteString(w, "\n")
				err = check(err, nerr)
				nerr = addindent(w, indent)
				err = check(err, nerr)
				_, nerr = io.WriteString(w, "}")
				err = check(err, nerr)
			default:
				return ErrUnsupportedType
			}
			_, nerr = io.WriteString(w, "\n")
			err = check(err, nerr)
		}
	default:
		return ErrUnsupportedType
	}

	return
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

func check(err, nerr error) error {
	if err != nil {
		return nerr
	}
	return err
}
