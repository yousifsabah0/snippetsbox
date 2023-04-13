package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestShowSnippet(t *testing.T) {
	app := NewTestApplication(t)
	ts := NewTestServer(t, app.Routes())
	defer ts.Close()

	tests := []struct {
		name     string
		path     string
		wantCode int
		wantBody []byte
	}{
		{"valid-id", "/snippets/1", http.StatusOK, []byte("Mocked Content")},
		{"non-existing-id", "/snippets/2", http.StatusNotFound, nil},
		{"negative-id", "/snippets/-1", http.StatusNotFound, nil},
		{"decimal-id", "/snippets/1.23", http.StatusNotFound, nil},
		{"string-id", "/snippets/foo", http.StatusNotFound, nil},
		{"empty-id", "/snippets/", http.StatusNotFound, nil},
		{"trailing-slash", "/snippets/1/", http.StatusNotFound, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.Get(t, test.path)

			if code != test.wantCode {
				t.Errorf("want %d; got %d", test.wantCode, code)
			}
			if !bytes.Contains(body, test.wantBody) {
				t.Errorf("want body to contain %q", test.wantBody)
			}
		})
	}
}
