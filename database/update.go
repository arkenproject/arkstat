package database

import (
	"database/sql"
	"log"
)

// Update changes a file's status in the database.
func Update(db *sql.DB, entry Node) {
	stmt, err := db.Prepare(
		`UPDATE nodes SET
			username = ?,
			email = ?,
			total_space = ?,
			used_space = ?,
			last_seen = datetime('now')
			WHERE id = ?;`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		entry.Username,
		entry.Email,
		entry.TotalSpace,
		entry.UsedSpace,
		entry.ID)

	if err != nil {
		log.Fatal(err)
	}
}
