package apiutils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// WriteError writes an error to the response
func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	if e, ok := err.(Error); ok {
		w.WriteHeader(e.StatusCode())
		bs, marshalErr := json.Marshal(errorResponse{
			Status: e.StatusCode(),
			Error:  e.Error(),
		})
		if marshalErr != nil {
			WriteError(w, marshalErr)
		}
		w.Write(bs)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	bs, marshalErr := json.Marshal(errorResponse{
		Status: http.StatusInternalServerError,
		Error:  http.StatusText(http.StatusInternalServerError),
	})
	if marshalErr != nil {
		WriteError(w, marshalErr)
	}
	w.Write(bs)
	return
}
