package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ramirezra/golang-web/section18/00-PostgresRefactor/books"
)

func main() {
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/create", books.CreateForm)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.UpdateForm)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}
