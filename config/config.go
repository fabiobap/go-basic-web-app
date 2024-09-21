package config

import "html/template"

//holds application cache
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
