package tasks

import (
	"log"
	"os"
	"time"
)

// MailGun Config Struct
type mailconf struct {
	Domain     string
	PrivateKey string
	Sender     string
	Setup      bool
}

var (
	config mailconf
)

// Initialize MailGun Config From ENV variables.
func init() {
	// Set initial config found value
	config.Setup = true
	// Check for MailGun Domain
	evVal, evExists := os.LookupEnv("STATS_MAIL_DOMAIN")
	if evExists {
		config.Domain = evVal
	} else {
		config.Setup = false
		log.Println("WARNING: MailGun Domain Not Found")
	}
	// Check for MailGun Private Key
	evVal, evExists = os.LookupEnv("STATS_MAIL_PRIVATEKEY")
	if evExists {
		config.PrivateKey = evVal
	} else {
		config.Setup = false
		log.Println("WARNING: MailGun Private Key Not Found")
	}
	// Check for MailGun Private Key
	evVal, evExists = os.LookupEnv("STATS_MAIL_SENDER")
	if evExists {
		config.Sender = evVal
	} else {
		config.Setup = false
		log.Println("WARNING: Sender Email Not Found")
	}
}

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
