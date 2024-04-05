package config

import (
	"dshusdock/tw_prac1/internal/constants"
	"log"
	"text/template"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	ViewCache     map[string]constants.ViewInteface
}
