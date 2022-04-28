package helpers

import (
	"encoding/json"
)

//string to type
func StringToType(s string, t interface{}) error {
	b := []byte(s)
	return json.Unmarshal(b, t)
}
