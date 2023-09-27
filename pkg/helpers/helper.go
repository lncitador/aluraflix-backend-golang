package helpers

import (
	"bytes"
	"encoding/json"
)

type Body map[string]any

func (b Body) MappedJson() *bytes.Buffer {
	body, _ := json.Marshal(b)

	return bytes.NewBuffer(body)
}

func MappedJson(data any) *bytes.Buffer {
	body, _ := json.Marshal(data)

	return bytes.NewBuffer(body)
}
