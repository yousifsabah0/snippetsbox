package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestApplication(t *testing.T) *Application {
	return &Application{
		ErrorLogger: log.New(io.Discard, "", 0),
		InfoLogger:  log.New(io.Discard, "", 0),
	}
}

type TestServer struct {
	*httptest.Server
}

func NewTestServer(t *testing.T, h http.Handler) *TestServer {
	ts := httptest.NewTLSServer(h)
	return &TestServer{ts}
}

func (ts *TestServer) Get(t *testing.T, path string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + path)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}
