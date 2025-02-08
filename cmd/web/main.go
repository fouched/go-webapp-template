package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/driver"
	"github.com/fouched/go-webapp-template/internal/handlers"
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/jaswdr/faker/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var app config.App
var session *scs.SessionManager

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// close database connection pool after application has stopped
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

	// define complex types that will be stored in the session
	//gob.Register(models.FooBar{})
	//gob.Register(map[string]int{})

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
	app.DB = db

	// seed the database ?
	//seed(db.SQL)

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// we can use persistent storage iso cookies for session data, this allows us to
	// restart the server without users losing the login / session information
	// https://github.com/alexedwards/scs has various options available
	//session.Store = pgxstore.New(db)
	app.Session = session

	// set up handlers and template rendering
	hc := handlers.NewHandlerConfig(&app)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)

	return db, nil
}

func seed(db *sql.DB) {
	fmt.Println("Start Seeding database")
	fake := faker.New()

	for i := 0; i < 100; i++ {
		stmt := `INSERT INTO customer (
                      customer_name, tel, email, address_1, address_2, address_3, post_code, created_at, updated_at)
    			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

		company := strings.Split(fake.Company().Name(), ",")[0]
		company = strings.Split(company, "-")[0]

		db.Exec(stmt,
			company,
			fake.Phone().E164Number(),
			"info@"+strings.ReplaceAll(company, " ", "")+".com",
			fake.Address().BuildingNumber()+" "+fake.Address().StreetName(),
			fake.Address().City(),
			fake.Address().State(),
			fake.Address().PostCode(),
			time.Now(),
			time.Now())
	}

	fmt.Println("End Seeding database")
}
