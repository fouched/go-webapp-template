package config

import (
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
	TemplateCache map[string]*template.Template
}
