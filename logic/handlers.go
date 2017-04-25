package logic

import (
	"net/http"

	"github.com/ramirezra/inv/config"
)

// Index exported
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	leads, err := AllLeads()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	config.Views.ExecuteTemplate(w, "index.gohtml", leads)
}

// Cards exported
func Cards(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	leads, err := AllLeads()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	config.Views.ExecuteTemplate(w, "cards.gohtml", leads)
}

// Tables exported
func Tables(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	leads, err := AllLeads()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	config.Views.ExecuteTemplate(w, "cards.gohtml", leads)
}
