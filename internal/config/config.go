package config

import (
	"html/template"
	"log"
)

type App struct {
	DSN           string
	Addr          string
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	TemplateCache map[string]*template.Template
}
