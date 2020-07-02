package tasks

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/arkenproject/arkstat/web"

	"github.com/arkenproject/arkstat/database"
)

// Start begins running a series of short tasks to update the website information every 2 minutes.
func Start() {
	for {
		db, err := database.Open(database.DatabaseLocation)
		if err != nil {
			log.Fatal(err)
		}

		// Calculates the total size of the pool from reporting nodes.
		total, used, err := database.GetPoolSize(db)
		if err != nil &&
			!strings.HasSuffix(err.Error(), "converting NULL to float64 is unsupported") {
			log.Fatal(err)
		}
		nodes, err := database.GetNodesReporting(db)
		if err != nil {
			log.Fatal(err)
		}
		// Set webpage values from Database.
		web.PageValues.TotalSpace = fmt.Sprintf("%f", float64(total)/float64(1000))
		web.PageValues.UsedSpace = fmt.Sprintf("%f", float64(used)/float64(1000))
		web.PageValues.ActiveNodes = nodes

		// Poll Database every two minutes.
		time.Sleep(2 * time.Minute)
	}
}
