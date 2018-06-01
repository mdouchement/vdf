#!/bin/bash

# ./peg2go.sh

cat > grammar.go <<EOF
package vdfp

// generated by peg2go shell script

import (
	"encoding/base64"
)

var grammar string

func getGrammar() (string, error) {
	if grammar != "" {
		return grammar, nil
	}

	b64 := "$(base64 grammar.peg)"

	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}

	grammar = string(data)
	return grammar, err
}
EOF