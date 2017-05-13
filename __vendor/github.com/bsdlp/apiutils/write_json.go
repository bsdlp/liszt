package apiutils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON is a helper that writes json to the response
func WriteJSON(w http.ResponseWriter, v interface{}) (err error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		return
	}
	return
}
