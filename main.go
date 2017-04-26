package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ramirezra/inv/logic"
)

func main() {
	http.HandleFunc("/", logic.Index)
	http.HandleFunc("/cards", logic.Cards)
	http.HandleFunc("/tables", logic.Tables)
	http.HandleFunc("/create", logic.CreateForm)
	http.HandleFunc("/create/process", logic.CreateProcess)
	http.HandleFunc("/update", logic.UpdateForm)
	http.HandleFunc("/update/process", logic.UpdateProcess)
	http.HandleFunc("/delete/process", logic.DeleteProcess)

	http.Handle("/views/css/", http.StripPrefix("/views/css", http.FileServer(http.Dir("./views/css"))))
	http.ListenAndServe(":8080", nil)
}
