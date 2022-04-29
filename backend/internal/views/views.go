package views

import (
	"embed"
	"html/template"
	"strings"
)

//go:embed *
var files embed.FS

var (
	layout     = "layout.gohtml"
	layoutHome = "layout_home.gohtml"
	homepage   = parse(layoutHome, "homepage/main.gohtml")

	templateFunctions = template.FuncMap{
		"concat": func(sep string, s ...string) string {
			return strings.Join(s, sep)
		},
		"split": strings.Split,
	}
)

func parse(layout, file string) *template.Template {
	return template.Must(
		template.New(layout).Funcs(templateFunctions).ParseFS(files, layout, file),
	)
}
