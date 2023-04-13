package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/yousifsabah0/snippetsbox/pkg/database/models/mocks"
)

func NewTestApplication(t *testing.T) *Application {
	templateCache, err := NewTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

	session := sessions.New([]byte("1937193nahda"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	return &Application{
		ErrorLogger:   log.New(io.Discard, "", 0),
		InfoLogger:    log.New(io.Discard, "", 0),
		Session:       session,
		Snippets:      &mocks.SnippetModel{},
		TemplateCache: templateCache,
		Users:         &mocks.UserModel{},
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
