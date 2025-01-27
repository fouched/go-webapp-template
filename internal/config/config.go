package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-webapp-template/internal/driver"
	"html/template"
	"log"
)

type App struct {
	DSN           string
	Addr          string
	DB            *driver.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}
