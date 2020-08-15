package database

import (
	"database/sql"
	"errors"
	"log"
)

// Get searches for and returns a the coorisponding entry from the
// database if the entry exists.
func Get(db *sql.DB, id string) (result Node, err error) {
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query("SELECT * FROM nodes WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if !row.Next() {
		return result, errors.New("entry not found")
	}
	err = row.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.TotalSpace,
		&result.UsedSpace,
		&result.LastSeen)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

// GetAllOld returns all of the node entries in the database in a channel.
func GetAllOld(db *sql.DB, output chan Node) {
	err := db.Ping()
	if err != nil {
		close(output)
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, email FROM nodes WHERE last_seen < datetime('now', '-1 day')")
	if err != nil {
		close(output)
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var node Node
		err = rows.Scan(
			&node.ID,
			&node.Email)

		if err != nil {
			close(output)
			log.Fatal(err)
		}
		output <- node
	}

	close(output)
}
