package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ramirezra/golang-web/section18/00-PostgresRefactor/config"
)

// Book exported
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks exported
func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, err
}

// OneBook exported to handlers.go
func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. BadRequest")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// PutBook exported to handlers.go
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. BadRequest")
	}

	// convert form values
	float, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("400. BadRequest")
	}
	bk.Price = float32(float)

	// insert values
	_, err = config.DB.Exec("INSERT INTO books (isbn,title,author,price) VALUES ($1,$2,$3,$4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())

	}
	return bk, nil
}

// UpdateBook exported to handlers.go
func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("406. Not Acceptable. Fields can't be blank")
	}

	// convert form values
	float, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(float)

	// insert values
	_, err = config.DB.Exec("UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// DeleteBook exported to handlers.go
func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	// delete bookstore
	_, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
