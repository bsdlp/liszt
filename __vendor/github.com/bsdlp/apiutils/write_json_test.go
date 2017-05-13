package apiutils

import (
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	type testDataStruct struct {
		A string `json:"a"`
		B bool   `json:"b"`
		C int    `json:"c"`
	}
	testData := testDataStruct{
		A: "hi",
		B: true,
		C: 420,
	}
	testJSON := `{"a":"hi","b":true,"c":420}`
	t.Run("writes the correct content type header", func(t *testing.T) {
		w := httptest.NewRecorder()
		WriteJSON(w, testData)
		contentType := w.HeaderMap.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("expected application/json as value for Content-Type, got %s", contentType)
		}
	})
	t.Run("writes correct data", func(t *testing.T) {
		w := httptest.NewRecorder()
		WriteJSON(w, testData)
		writtenJSON := w.Body.String()
		if writtenJSON != testJSON {
			t.Errorf("expected %s, got %s", testJSON, writtenJSON)
		}
	})
}
