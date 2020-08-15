package database

import (
	"database/sql"
)

// Remove deletes an node from the database.
func Remove(tx *sql.Tx, id string) (err error) {
	stmt, err := tx.Prepare("DELETE FROM nodes WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
