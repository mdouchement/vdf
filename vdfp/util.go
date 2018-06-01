package vdfp

const (
	tRoot  = "Root"
	tNode  = "Node"
	tKVs   = "KVs"
	tKey   = "Key"
	tValue = "Value"
)

func unQuote(s string) string {
	return s[1 : len(s)-1]
}
