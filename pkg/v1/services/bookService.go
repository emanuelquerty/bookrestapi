package services

import (
	"bookrestapi/pkg/v1/models"
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

func GetAllBooks(ctx context.Context, db *sql.DB) ([]models.Book, error) {
	query := "SELECT * FROM book"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("books %v", err)
	}

	books := []models.Book{}
	for rows.Next() {
		book := models.Book{}
		s := reflect.ValueOf(&book).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		if err != nil {
			return nil, fmt.Errorf("no books found. Error: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBook(ctx context.Context, db *sql.DB, id string) (models.Book, error) {
	var book models.Book

	query := "SELECT * FROM book WHERE book_id = $1"
	row := db.QueryRowContext(ctx, query, id)

	if err := row.Scan(&book.BookID, &book.Title, &book.PublishedYear, &book.AuthorID); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("booksById %s: no such book", id)
		}
		return book, fmt.Errorf("booksById %s: %v", id, err)
	}
	return book, nil
}

func PostBook(ctx context.Context, db *sql.DB, book models.Book) (int64, error) {
	var id int64
	query := "INSERT INTO book(title, published_year, author_id) VALUES($1, $2, $3) RETURNING book_id"
	err := db.QueryRowContext(ctx, query, book.Title, book.PublishedYear, book.AuthorID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("add book error: %v", err)
	}
	return id, nil
}

func UpdateBook(ctx context.Context, db *sql.DB, book models.Book, id string) (models.Book, error) {
	query := `
	UPDATE book
	SET title = $1, published_year = $2, author_id = $3
	WHERE book_id = $4
	RETURNING book_id, title, published_year, author_id`
	row := db.QueryRowContext(ctx, query, book.Title, book.PublishedYear, book.AuthorID, id)
	updatedBook := models.Book{}
	err := row.Scan(&updatedBook.BookID, &updatedBook.Title, &updatedBook.PublishedYear, &updatedBook.AuthorID)
	if err != nil {
		return updatedBook, fmt.Errorf("update book error: %v", err)
	}
	return updatedBook, nil
}

func DeleteBook(ctx context.Context, db *sql.DB, id string) ([]models.Book, error) {
	query := `DELETE FROM book WHERE book_id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("delete book error: %v", err)
	}
	query = `SELECT * FROM book`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("books: %v", err)
	}
	books := []models.Book{}
	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.BookID, &book.Title, &book.PublishedYear, &book.AuthorID)
		if err != nil {
			return nil, fmt.Errorf("books: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}
