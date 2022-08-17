package graphql

import (
	"bytes"
	"encoding/json"
	"io"
)

type V map[string]any

type RequestBody struct {
	Query     string `json:"query"`
	Variables V      `json:"variables"`
}

func CreateRequest(query string, variables V) (io.Reader, error) {
	reqBody := RequestBody{
		Query:     query,
		Variables: variables,
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(reqBody); err != nil {
		return nil, err
	}
	return b, nil
}
