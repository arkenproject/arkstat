package database

import (
	"database/sql"
	"errors"
)

// GetPoolSize return the sum the nodes total space and used space.
func GetPoolSize(db *sql.DB) (total float64, used float64, err error) {
	err = db.Ping()
	if err != nil {
		return total, used, err
	}
	// Get total pool size from sum of nodes reported values.
	row, err := db.Query("SELECT SUM(total_space) FROM nodes")
	if err != nil {
		return total, used, err
	}
	defer row.Close()
	if !row.Next() {
		return total, used, errors.New("sum not found")
	}
	err = row.Scan(&total)
	if err != nil {
		return total, used, err
	}

	// Get used pool size from sum of nodes reported values.
	row, err = db.Query("SELECT SUM(used_space) FROM nodes")
	if err != nil {
		return total, used, err
	}
	defer row.Close()
	if !row.Next() {
		return total, used, errors.New("sum not found")
	}
	err = row.Scan(&used)
	if err != nil {
		return total, used, err
	}
	return total, used, nil
}

// GetNodesReporting calculates the number of nodes reporting to the database.
func GetNodesReporting(db *sql.DB) (nodes int, err error) {
	err = db.Ping()
	if err != nil {
		return -1, err
	}
	// Get total number of nodes reporting.
	row, err := db.Query("SELECT COUNT(id) FROM nodes")
	if err != nil {
		return -1, err
	}
	defer row.Close()
	if !row.Next() {
		return nodes, errors.New("no nodes found")
	}
	err = row.Scan(&nodes)
	if err != nil {
		return -1, err
	}
	return nodes, nil
}
