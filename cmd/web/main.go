package main

import (
	"flag"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/driver"
	"github.com/fouched/go-webapp-template/internal/handlers"
	"github.com/fouched/go-webapp-template/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
)

var app config.App

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// close database connection after application has stopped
	defer db.SQL.Close()

	mux := routes()
	srv := &http.Server{
		Addr:    app.Addr,
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run() (*driver.DB, error) {

	// create the template cache
	app.TemplateCache = make(map[string]*template.Template)

	// read config
	flag.StringVar(&app.Addr, "addr", ":9080", "Server addr to listen on")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=password dbname=webapp sslmode=disable", "DSN (Data Source Name)")

	// init loggers
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// connect to db
	db, err := driver.ConnectSQL(app.DSN)
	if err != nil {
		return nil, err
	}
	app.InfoLog.Println("Connected to DB")

	// set up handlers and template rendering
	hc := handlers.NewHandlerConfig(&app, db)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)

	return db, nil
}
