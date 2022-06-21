package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbSource = ""
)

func Connect() (context.Context, *sql.DB) {
	ctx := context.Background()
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("CONN ERR: ", err)
	}

	// Verify that a connection can be made
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("VERIF CONN ERR: ", err)
	}
	return ctx, db
}
