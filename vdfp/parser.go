package vdfp

import (
	"fmt"

	peg "github.com/yhirose/go-peg"
)

// Parse returns the unmarshalled Valve Data Format blob.
func Parse(blob string) (map[string]interface{}, error) {
	g, err := getGrammar()
	if err != nil {
		return nil, fmt.Errorf("Error during grammar unblob: %s", err)
	}

	parser, err := peg.NewParser(g)
	if err != nil {
		return nil, fmt.Errorf("Error in parser compilation: %s", err)
	}
	setEvaluationRules(parser.Grammar)

	val, err := parser.ParseAndGetValue(blob, nil)
	if err != nil {
		return nil, fmt.Errorf("Error during blob parsing: %s", err)
	}

	return val.(map[string]interface{}), nil
}
