package vdfp

import (
	"fmt"

	peg "github.com/yhirose/go-peg"
)

// A Node is a part of the AST.
type Node struct {
	// N is the name of the node.
	N string
	// V is the value of the node.
	V peg.Any
}

// String implements Stringer.
func (n *Node) String() string {
	if n.V == nil {
		return fmt.Sprintf("<%s>", n.N)
	}
	return fmt.Sprintf("<%s>(%v)", n.N, n.V)
}

func setEvaluationRules(g map[string]*peg.Rule) {
	// ==================== //
	// Root                 //
	// ==================== //

	g[tRoot].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		root := make(map[string]interface{})

		for _, val := range v.Vs {
			n := val.(*Node)

			switch n.N {
			case tKVs:
				for k, v := range n.V.(map[string]interface{}) {
					root[k] = v // String
				}
			case tNode:
				for k, v := range n.V.(map[string]interface{}) {
					root[k] = v // Object
				}
			}
		}

		return root, nil
	}

	// ==================== //
	// Node                 //
	// ==================== //

	g[tNode].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		var key string
		var value = make(map[string]interface{})

		for _, val := range v.Vs {
			n := val.(*Node)

			switch n.N {
			case tKey:
				key = n.V.(string)
			case tKVs:
				for k, v := range n.V.(map[string]interface{}) {
					value[k] = v // String
				}
			case tNode:
				for k, v := range n.V.(map[string]interface{}) {
					value[k] = v // Object
				}
			}
		}

		return &Node{N: tKVs, V: map[string]interface{}{key: value}}, nil
	}

	// ==================== //
	// Key/Value            //
	// ==================== //

	g[tKVs].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		var key string
		var value interface{}

		for _, val := range v.Vs {
			n := val.(*Node)

			switch n.N {
			case tKey:
				key = n.V.(string)
			case tValue:
				value = n.V
			}
		}

		return &Node{N: tKVs, V: map[string]interface{}{key: value}}, nil
	}

	g[tKey].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return &Node{N: tKey, V: unQuote(v.Token())}, nil
	}

	g[tValue].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return &Node{N: tValue, V: unQuote(v.Token())}, nil
	}

	// ==================== //
	// Basics               //
	// ==================== //

	g["T"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return &Node{N: "T"}, nil
	}

	g["N"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return &Node{N: "N"}, nil
	}

	g["S"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return &Node{N: "S"}, nil
	}
}
