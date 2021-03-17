package jsondecoder

import (
	"encoding/json"
	"fmt"
	"io"
)

type JsonDecoder interface {
	Decode(r io.Reader) (map[string]string, error)
}

type MyJsonDecoder struct {
}

func (d MyJsonDecoder) Decode(r io.Reader) (map[string]string, error) {

	var m map[string]interface{}
	err := json.NewDecoder(r).Decode(&m)

	m2 := make(map[string]string)
	// TODO: good solution?
	for k := range m {
		m2[k] = fmt.Sprintf("%v", m[k])
	}

	return m2, err
}
