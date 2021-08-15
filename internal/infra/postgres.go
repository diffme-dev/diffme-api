package infra

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgresConnection() *sql.DB {
	db, err := sql.Open("mongo", "string...")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}
