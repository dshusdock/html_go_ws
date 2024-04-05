package render

import (
	"html/template"
	"net/http"
)


var files = []string{
	"./ui/html/pages/base.tmpl",
	"./ui/html/pages/layout.tmpl",
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, d any) {
	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "base", d)
}

