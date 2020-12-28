package config

import (
	"log"
	"os"
	"strings"
)

// MailGun Config Struct
type mailconf struct {
	Domain     string
	PrivateKey string
	Sender     string
	Setup      bool
}

var (
	// Mail is the mail config for the Arkstat instance.
	Mail mailconf
)

// Initialize MailGun Config From ENV variables.
func init() {
	// Set initial config found value
	Mail.Setup = true
	// Check for MailGun Domain
	evVal, evExists := os.LookupEnv("STATS_MAIL_DOMAIN")
	if evExists {
		Mail.Domain = evVal
	} else {
		Mail.Setup = false
		log.Println("WARNING: MailGun Domain Not Found")
	}
	// Check for MailGun Private Key
	evVal, evExists = os.LookupEnv("STATS_MAIL_PRIVATEKEY")
	if evExists {
		Mail.PrivateKey = evVal
	} else {
		Mail.Setup = false
		log.Println("WARNING: MailGun Private Key Not Found")
	}
	// Check for MailGun Private Key
	evVal, evExists = os.LookupEnv("STATS_MAIL_SENDER")
	if evExists {
		Mail.Sender = evVal
	} else {
		Mail.Setup = false
		log.Println("WARNING: Sender Email Not Found")
	}
}

// CleanEmail clean the email address input.
func CleanEmail(email string) string {
	return strings.TrimSuffix(strings.TrimPrefix(email, `"`), `"`)
}
