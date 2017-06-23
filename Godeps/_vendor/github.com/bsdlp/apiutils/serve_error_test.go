package apiutils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	type testData struct {
		statusCode int
		bodyString string
		err        error
	}
	cases := []testData{
		{http.StatusInternalServerError, `{"status":500,"error":"Internal Server Error"}`, errors.New("")},
		{http.StatusInternalServerError, `{"status":500,"error":"Internal Server Error"}`, errors.New("blah")},
		{http.StatusBadGateway, `{"status":502,"error":"Bad Gateway"}`, NewError(http.StatusBadGateway, "")},
		{http.StatusTeapot, `{"status":418,"error":"lulz"}`, NewError(http.StatusTeapot, "lulz")},
	}
	for _, testCase := range cases {
		w := httptest.NewRecorder()
		WriteError(w, testCase.err)
		t.Run("writes the right http status code", func(t *testing.T) {
			if w.Code != testCase.statusCode {
				t.Errorf("expected %d, got %d", testCase.statusCode, w.Code)
			}
		})
		t.Run("writes the right body", func(t *testing.T) {
			bodyString := w.Body.String()
			if bodyString != testCase.bodyString {
				t.Errorf("expected %s, got %s", testCase.bodyString, bodyString)
			}
		})
	}
}
