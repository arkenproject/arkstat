package tasks

import (
	"log"
	"time"
)

// Start begins running a series of short tasks to update the website information every 2 minutes.
func Start() {
	for {
		err := removeMissing()
		if err != nil {
			log.Println(err)
		}

		// Update the values of the site without querying the database every page load.
		err = updateSite()
		if err != nil {
			log.Println(err)
		}

		// Poll Database every two minutes.
		time.Sleep(2 * time.Minute)
	}
}
