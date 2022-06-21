package main

import (
	"bookrestapi/pkg/db"
	"bookrestapi/pkg/v1/routes"
	"log"
	"net/http"
)

func main() {
	ctx, db := db.Connect()
	http.HandleFunc("/api/v1/books", routes.HttpMethodHandler(ctx, db))
	http.HandleFunc("/api/v1/books/", routes.HttpMethodHandler(ctx, db))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ERR SERV LAUNCH", err)
	}
}
