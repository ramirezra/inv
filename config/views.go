package config

import "html/template"

// Views exported
var Views *template.Template

func init() {
	Views = template.Must(template.ParseGlob("views/templates/*.gohtml"))
}
