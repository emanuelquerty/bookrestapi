package routes

import (
	"bookrestapi/pkg/v1/controllers"
	"context"
	"database/sql"
	"net/http"
)

// func MakeHandler(fn func(http.ResponseWriter, *http.Request, context.Context, *sql.DB),
// 	ctx context.Context, db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fn(w, r, ctx, db)
// 	}
// }

func HandleResourceNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("The requested resource does not exist."))
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("The Method used is not allowed."))
}

func HttpMethodHandler(ctx context.Context, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.PostBook(w, r, ctx, db)
		case http.MethodGet:
			if r.URL.Path == "/api/v1/books" {
				controllers.GetAllBooks(w, r, ctx, db)
				return
			}
			controllers.GetBook(w, r, ctx, db)
		case http.MethodPut:
			controllers.UpdateBook(w, r, ctx, db)
		case http.MethodDelete:
			controllers.DeleteBook(w, r, ctx, db)
		default:
			HandleMethodNotAllowed(w, r)
		}
	}
}
