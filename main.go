package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Lead exported
type Lead struct {
	ID      string
	Status  string
	Contact string
	Sales   string
	Value   string
}

var dbLeads = map[int]Lead{}

var tpl *template.Template

func init() {
	tpl = template.
		Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	// var table []string

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1

	record, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var leads []Lead

	for _, item := range record {
		// ID, _ := strconv.ParseInt(item[1], 1, 1)
		// Status, _ := item[]
		// open, _ := strconv.ParseFloat(item[1], 64)
		// high, _ := strconv.ParseFloat(record[2], 64)
		// low, _ := strconv.ParseFloat(record[3], 64)
		// close, _ := strconv.ParseFloat(record[4], 64)
		// volume, _ := strconv.ParseFloat(record[5], 64)

		lead := Lead{ID: item[0], Status: item[1]}
		leads = append(leads, lead)
	}

	error := tpl.Execute(os.Stdout, leads)
	if err != nil {
		log.Fatalln(error)
	}
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", leads)
	})
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}
