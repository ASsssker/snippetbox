package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

var info = color.New(color.FgHiGreen).Sprint("INFO\t")
var errorL = color.New(color.FgRed).Sprint("ERROR\t")

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// addr is a pointer to command line option with the nme "addr".
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parsing command line options.
	flag.Parse()

	// Create a info logger.
	infoLog := log.New(os.Stdout, info, log.Ldate|log.Ltime)
	// Create a error log
	errorLog := log.New(os.Stderr, errorL, log.Ldate|log.Ltime|log.Lshortfile)
	
	// Initialize a new instance of application struct.
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	// Use ListenAndServe func to start a new web server.
	// We pass in tcp socket address and ServeMux router.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}