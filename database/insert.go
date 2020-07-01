package database

import (
	"database/sql"
	"log"
)

// Add inserts a Node entry into the database if it doesn't exist already.
func Add(db *sql.DB, input Node) (err error) {
	_, err = Get(db, input.ID)
	if err != nil {
		if err.Error() == "entry not found" {
			insert(db, input)
		} else {
			return err
		}
	}
	Update(db, input)
	return nil
}

// Insert adds a Node entry to the database.
func insert(db *sql.DB, entry Node) {
	// Ping to check that database connection still exists.
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(
		`INSERT INTO nodes(
			id,
			username,
			email,
			total_space,
			used_space,
			last_seen
		) VALUES(?,?,?,?,?,datetime('now'));`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		entry.ID,
		entry.Username,
		entry.Email,
		entry.TotalSpace,
		entry.UsedSpace)

	if err != nil {
		log.Fatal(err)
	}
}
