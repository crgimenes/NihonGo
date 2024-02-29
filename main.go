package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

const (
	connectionString = `file:nihongo.db?mode=rwc&_journal_mode=WAL&_busy_timeout=10000`
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := sqlx.Open("sqlite", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
