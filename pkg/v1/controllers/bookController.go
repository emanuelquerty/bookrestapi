package controllers

import (
	"bookrestapi/pkg/v1/models"
	"bookrestapi/pkg/v1/services"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request, ctx context.Context, db *sql.DB) {
	books, err := services.GetAllBooks(ctx, db)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	BookResponse := models.BookResponse{Status: "sucess", StatusCode: 200, Data: books}
	booksJson, err := json.Marshal(BookResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(booksJson)
}

func GetBook(w http.ResponseWriter, r *http.Request, ctx context.Context, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/api/v1/books/"):]
	book, err := services.GetBook(ctx, db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	BookResponse := models.BookResponse{Status: "sucess", StatusCode: 200, Data: []models.Book{book}}
	if err := json.NewEncoder(w).Encode(BookResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostBook(w http.ResponseWriter, r *http.Request, ctx context.Context, db *sql.DB) {
	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := services.PostBook(ctx, db, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := struct {
		StatusCode   int    `json:"status_code,omitempty"`
		Status       string `json:"status,omitempty"`
		LastInsertID int64  `json:"last_insert_id,omitempty"`
	}{
		http.StatusCreated,
		"Sucess",
		id,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request, ctx context.Context, db *sql.DB) {
	id := r.URL.Path[len("/api/v1/books/"):]
	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedBook, err := services.UpdateBook(ctx, db, book, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := models.BookResponse{
		Status:     "Success",
		StatusCode: 200,
		Data:       []models.Book{updatedBook},
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request, ctx context.Context, db *sql.DB) {
	id := r.URL.Path[len("/api/v1/books/"):]
	books, err := services.DeleteBook(ctx, db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := models.BookResponse{
		Status:     "Success",
		StatusCode: 200,
		Data:       books,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
