package utils

import (
	"encoding/json"
	"net/http"
	"io"
)

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}