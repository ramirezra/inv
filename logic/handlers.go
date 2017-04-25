package logic

import (
	"net/http"

	"github.com/ramirezra/inv/config"
)

// Index exported
func Index(w http.ResponseWriter, r *http.Request) {
	config.Views.ExecuteTemplate(w, "index.gohtml", nil)
}

// Cards exported
func Cards(w http.ResponseWriter, r *http.Request) {
	config.Views.ExecuteTemplate(w, "cards.gohtml", data)
}

// Tables exported
func Tables(w http.ResponseWriter, r *http.Request) {
	config.Views.ExecuteTemplate(w, "cards.gohtml", data)
}
