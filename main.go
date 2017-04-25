package main

import (
	"net/http"

	"github.com/ramirezra/inv/logic"
)

func main() {
	http.Handle("/", logic.Index)
	http.Handle("/cards", logic.Cards)
	http.Handle("/tables", logic.Tables)

	http.ListenAndServe(":8080", nil)

}
