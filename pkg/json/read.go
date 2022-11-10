package json

import (
	"encoding/json"
	"io"
)

func Read[T interface{}](reader io.ReadCloser) (*T, error) {
	var jsonData T

	if err := json.NewDecoder(reader).Decode(&jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}
