package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/yousifsabah0/snippetsbox/pkg/database"
	"github.com/yousifsabah0/snippetsbox/pkg/database/models/mysql"
)

type Application struct {
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Snippets      *mysql.SnippetModel
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

	app := &Application{
		InfoLogger:    infoLog,
		ErrorLogger:   errorLog,
		Snippets:      &mysql.SnippetModel{Db: db},
		TemplateCache: tc,
		Session:       session,
	}

	// Create & Start the web server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("up and running.. port: %s", *addr)
	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
