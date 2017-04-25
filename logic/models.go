package logic

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Lead exported
type Lead struct {
	ID      string
	Status  string
	Contact string
	Sales   string
	Value   float64
}

var dbLeads = map[int]Lead{}

// GetData exported
func GetData() {
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

	for i, item := range record {
		if i == 0 {
			continue
		}
		// open, _ := strconv.ParseFloat(item[1], 64)
		value, _ := strconv.ParseFloat(item[4], 64)
		lead := Lead{ID: item[0], Status: item[1], Contact: item[2], Sales: item[3], Value: value}
		leads = append(leads, lead)
		fmt.Println(leads)
	}

	error := Views.Execute(os.Stdout, leads)
	if err != nil {
		log.Fatalln(error)
	}
}
