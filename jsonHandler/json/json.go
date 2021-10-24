package jsonhandler

import (
	"encoding/json"

	jsonhandler "github.com/Kasparund/Go-Action-Test-Overload/jsonHandler"
)

type jsonLib struct{}

func NewJSONHandler() jsonhandler.JSONHandler {
	return &jsonLib{}
}

func (j jsonLib) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (j jsonLib) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
