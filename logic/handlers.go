package logic

import (
	"database/sql"
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
	config.Views.ExecuteTemplate(w, "table.gohtml", leads)
}

// CreateForm exported
func CreateForm(w http.ResponseWriter, r *http.Request) {
	config.Views.ExecuteTemplate(w, "create.gohtml", nil)
}

// CreateProcess exported
func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	lead, err := CreateLead(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
	}

	// confirm Creation
	config.Views.ExecuteTemplate(w, "created.gohtml", lead)
}

// UpdateForm exported
func UpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM leads WHERE id=$1", id)

	lead := Lead{}
	err := row.Scan(&lead.ID, &lead.Status, &lead.Contact, &lead.Sales, &lead.Value)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	config.Views.ExecuteTemplate(w, "update.gohtml", lead)
}

// UpdateProcess exported
func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	lead, err := UpdateLead(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
	}
	// confirm update
	config.Views.ExecuteTemplate(w, "updated.gohtml", lead)
}

// DeleteProcess exported
func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	err := DeleteLead(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
