package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edwlarkey/ril/pkg/bolt"
	"github.com/edwlarkey/ril/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type contextKey string

var contextKeyUser = contextKey("user")

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	db       interface {
		Connect(string) error
		Close()
		InsertArticle(string, string, string) (int, error)
		GetArticle(int) (*models.Article, error)
		LatestArticles() ([]*models.Article, error)
		InsertUser(string, string, string) error
		AuthenticateUser(string, string) (int, error)
		GetUser(int) (*models.User, error)
	}
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "ril.db", "Data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		templateCache: templateCache,
		db:            &bolt.DB{},
	}

	// Connect to the DB
	err = app.db.Connect(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Defer closing our DB connection pool
	defer app.db.Close()

	// Set up http server, including app routes
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the http server
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
