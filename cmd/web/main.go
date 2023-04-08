package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":8080", "HTTP network port")

	flag.Parse()

	// Create custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Llongfile|log.Ldate)

	app := &Application{
		InfoLogger:  infoLog,
		ErrorLogger: errorLog,
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
