package helpers

import (
	"encoding/json"
)

//string to type
func GenericToType[E any](s string, t E) error {
	b := []byte(s)
	return json.Unmarshal(b, t)
}

func StringToType(s string, t interface{}) error {
	b := []byte(s)
	return json.Unmarshal(b, t)
}
