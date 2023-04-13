package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := NewTestApplication(t)
	ts := NewTestServer(t, app.Routes())
	defer ts.Close()

	code, _, body := ts.Get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("got %v; want %v", code, http.StatusOK)
	}

	if string(body) != "ok" {
		t.Errorf("got %v; want %v", string(body), "ok")
	}
}
