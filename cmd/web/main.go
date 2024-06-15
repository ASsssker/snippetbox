package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"snippetbox/internal/models"

	"github.com/fatih/color"

	_ "github.com/go-sql-driver/mysql"
)

var info = color.New(color.FgHiGreen).Sprint("INFO\t")
var errorL = color.New(color.FgRed).Sprint("ERROR\t")

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// addr is a pointer to command line option with the name "addr".
	addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn is a pointer to command line option with database destination addr.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=True", "Database source name")
	// Parsing command line options.
	flag.Parse()

	// Create a info logger.
	infoLog := log.New(os.Stdout, info, log.Ldate|log.Ltime)
	// Create a error logger.
	errorLog := log.New(os.Stderr, errorL, log.Ldate|log.Ltime|log.Lshortfile)

	// Get db connection pool.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new instance of application struct.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	// Use ListenAndServe func to start a new web server.
	// We pass in tcp socket address and ServeMux router.
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns
// sql.DB connection pool.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
