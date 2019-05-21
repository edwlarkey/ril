package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/edwlarkey/ril/pkg/config"
	"github.com/edwlarkey/ril/pkg/models"
	"github.com/edwlarkey/ril/pkg/mysql"
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
	configFile := flag.String("config", "./config.toml", "Config file")
	flag.Parse()

	var conf config.Config
	if _, err := toml.DecodeFile(*configFile, &conf); err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Port, conf.Database.Database)
	addr := fmt.Sprintf("%s:%s", conf.Web.IP, conf.Web.Port)

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(conf.Web.Secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		db:            &mysql.DB{},
		templateCache: templateCache,
	}

	// Connect to the DB
	err = app.db.Connect(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Defer closing our DB connection pool
	defer app.db.Close()

	// Set up http server, including app routes
	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the http server
	infoLog.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
