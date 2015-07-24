package util

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("\n"))

	return err
}
