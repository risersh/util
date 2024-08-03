// Package json provides a wrapper around the standard library json package.
package json

import "encoding/json"

// MarshalStrict marshals the given value to JSON, panicking if the value cannot be marshaled.
func MarshalStrict(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
