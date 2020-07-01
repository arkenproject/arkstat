package database

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver for database interaction.
)

// Node defines the columns of data stored in the database.
type Node struct {
	ID         string
	Username   string
	Email      string
	TotalSpace int
	UsedSpace  int
	LastSeen   time.Time
}

// DatabaseLocation is the path of the Stats Database
var DatabaseLocation string

// Determine location of database.
func init() {
	evVal, evExists := os.LookupEnv("STATS_DATABASE")
	if evExists {
		DatabaseLocation = evVal
	} else {
		DatabaseLocation = "./stats.db"
	}

}

// Open connects and returns a pointer to a database object.
func Open(path string) (db *sql.DB, err error) {

	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create the nodes table if is doesn't already exist.
	// This will also create the database if it doesn't exist.
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS nodes(
			id TEXT,
			username TEXT,
			email TEXT,
			total_space INT,
			used_space TEXT,
			last_seen DATETIME,
			
			PRIMARY KEY(id)
		);`)

	if err != nil {
		return nil, err
	}

	return db, nil
}
