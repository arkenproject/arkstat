package tasks

import (
	"fmt"
	"strings"

	"github.com/arken/arkstat/database"
	"github.com/arken/arkstat/web"
)

func updateSite() (err error) {
	db, err := database.Open(database.DatabaseLocation)
	if err != nil {
		return err
	}

	// Calculates the total size of the pool from reporting nodes.
	total, used, err := database.GetPoolSize(db)
	if err != nil &&
		!strings.HasSuffix(err.Error(), "converting NULL to float64 is unsupported") {
		return err
	}
	nodes, err := database.GetNodesReporting(db)
	if err != nil {
		return err
	}
	// Set webpage values from Database.
	web.PageValues.TotalSpace = fmt.Sprintf("%.2f", float64(total)/float64(1000))
	web.PageValues.UsedSpace = fmt.Sprintf("%.4f", float64(used)/float64(1000))
	web.PageValues.ActiveNodes = nodes

	return db.Close()
}
