package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrInvalidContentType = errors.New("Content-Type is not application/json")

func RequestJSON(r *http.Request, ptr any) error {
	// Check the content type
	if r.Header.Get("Content-Type") != "application/json" {
		err := ErrInvalidContentType
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(ptr); err != nil {
		return err
	}

	return nil
}
