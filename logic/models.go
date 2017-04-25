package logic

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ramirezra/inv/config"
)

// Lead exported
type Lead struct {
	ID      string
	Status  string
	Contact string
	Sales   string
	Value   float64
}

// AllLeads exported - Read (all) part of CRUD
func AllLeads() ([]Lead, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	leads := make([]Lead, 0)
	for rows.Next() {
		lead := Lead{}
		err := rows.Scan(&lead.ID, &lead.Status, &lead.Contact, &lead.Sales, &lead.Value)
		if err != nil {
			return nil, err
		}
		leads = append(leads, lead)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return leads, err
}

// OneLead exporte to handlers.go | Read part
func OneLead(r *http.Request) (Lead, error) {
	lead := Lead{}
	id := r.FormValue("id")
	if id == "" {
		return lead, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM leads WHERE id =$1", id)
	err := row.Scan(&lead.ID, &lead.Status, &lead.Contact, &lead.Sales, &lead.Value)
	if err != nil {
		return lead, err
	}
	return lead, nil
}

// CreateLead exported to handlers.go | Create (one) part of CRUD
func CreateLead(r *http.Request) (Lead, error) {
	// get form values
	lead := Lead{}
	lead.ID = r.FormValue("id")
	lead.Status = r.FormValue("title")
	lead.Contact = r.FormValue("contact")
	lead.Sales = r.FormValue("sales")
	v := r.FormValue("value")

	// validate form values
	if lead.ID == "" || lead.Status == "" || lead.Contact == "" || lead.Sales == "" || v == "" {
		return lead, errors.New("400. Bad Request")
	}
	// convert form values
	float, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return lead, errors.New("400. Bad Request")
	}
	lead.Value = float

	// insert values
	_, err = config.DB.Exec("INSERT INTO leads (id,status,contact,sales,value) VALUES ($1,$2,$3,$4,$5)", lead.ID, lead.Status, lead.Contact, lead.Sales, lead.Value)
	if err != nil {
		return lead, errors.New("500. Internal Server Error" + err.Error())
	}
	return lead, nil
}

// UpdateLead exported to handlers.go | Update part of CRUD
func UpdateLead(r *http.Request) (Lead, error) {
	// get form values
	lead := Lead{}
	lead.ID = r.FormValue("id")
	lead.Status = r.FormValue("title")
	lead.Contact = r.FormValue("contact")
	lead.Sales = r.FormValue("sales")
	v := r.FormValue("value")

	// validate form values
	if lead.ID == "" || lead.Status == "" || lead.Contact == "" || lead.Sales == "" || v == "" {
		return lead, errors.New("400. Bad Request")
	}
	// convert form values
	float, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return lead, errors.New("400. Bad Request")
	}
	lead.Value = float

	// Update values
	_, err = config.DB.Exec("UPDATE leads SET id=$1, status=$2, contact=$3, sales=$4, value=$5", lead.ID, lead.Status, lead.Contact, lead.Sales, lead.Value)
	if err != nil {
		return lead, err
	}
	return lead, nil
}

// DeleteLead exported to handlers.go | Delete part of CRUD
func DeleteLead(r *http.Request) error {
	id := r.FormValue("id")
	if id == "" {
		return errors.New("400. Bad Request")
	}

	// delete lead
	_, err := config.DB.Exec("DELETE FROM leads WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

// var dbLeads = map[int]Lead{}
//
// // GetData exported
// func GetData() {
// 	file, err := os.Open("data.csv")
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}
// 	defer file.Close()
//
// 	// var table []string
//
// 	reader := csv.NewReader(file)
//
// 	reader.FieldsPerRecord = -1
//
// 	record, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	var leads []Lead
//
// 	for i, item := range record {
// 		if i == 0 {
// 			continue
// 		}
// 		// open, _ := strconv.ParseFloat(item[1], 64)
// 		value, _ := strconv.ParseFloat(item[4], 64)
// 		lead := Lead{ID: item[0], Status: item[1], Contact: item[2], Sales: item[3], Value: value}
// 		leads = append(leads, lead)
// 		fmt.Println(leads)
// 	}
//
// 	error := Views.Execute(os.Stdout, leads)
// 	if err != nil {
// 		log.Fatalln(error)
// 	}
// }
