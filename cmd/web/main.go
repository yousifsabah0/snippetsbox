package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/yousifsabah0/snippetsbox/pkg/database"
	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
	"github.com/yousifsabah0/snippetsbox/pkg/database/models/mysql"
)

type ContextKey string

const (
	ContextKeyIsAuthenticated = ContextKey("IsAuthenticated")
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Snippets    interface {
		Insert(title, content, expires string) (int, error)
		Get(id int) (*models.Snippet, error)
		Latest() ([]*models.Snippet, error)
		Update(snippet *models.Snippet) (int, error)
		Delete(id int) error
	}
	Users interface {
		Insert(name, email, password string) error
		Authenticate(email, pass string) (int, error)
		Get(id int) (*models.User, error)
	}
	TemplateCache TemplateCache
	Session       *sessions.Session
}

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":8080", "HTTP network port")
	dsn := flag.String("dsr", "stark:1538@/snippetsbox?parseTime=true", "Database source name")
	secret := flag.String("session", "1937193nahda", "Session secret ket")

	flag.Parse()

	// Create custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Llongfile|log.Ldate)

	// Connect to database
	db, err := database.Open(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Template caching
	tc, err := NewTemplateCache("./web/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Set up session
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &Application{
		InfoLogger:    infoLog,
		ErrorLogger:   errorLog,
		Snippets:      &mysql.SnippetModel{Db: db},
		Users:         &mysql.UserModel{Db: db},
		TemplateCache: tc,
		Session:       session,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Create & Start the web server
	server := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.Routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("up and running.. port: %s", *addr)
	if err := server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"); err != nil {
		errorLog.Fatal(err)
	}
}
