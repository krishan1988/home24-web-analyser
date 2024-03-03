package root

import (
	"embed"
	"html/template"
)

//go:embed templates/*.tmpl
var embededTemplates embed.FS

func GetTemplates() *template.Template {
	return template.Must(template.ParseFS(embededTemplates, "templates/*.tmpl"))
}