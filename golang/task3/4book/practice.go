package _book

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  sql.NullFloat64 `db:"price"`
}

func Run(sqlx *sqlx.DB) {
	var books []Book
	err := sqlx.Select(&books, "select * from books")
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
