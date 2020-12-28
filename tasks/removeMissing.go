package tasks

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/arkenproject/arkstat/config"
	"github.com/arkenproject/arkstat/database"
	"github.com/mailgun/mailgun-go/v4"
)

// Remove nodes who haven't checked in, in a day from the database and send
// an email if provided.
func removeMissing() (err error) {
	output := make(chan database.Node)
	var mg *mailgun.MailgunImpl
	if config.Mail.Setup {
		mg = mailgun.NewMailgun(config.Mail.Domain, config.Mail.PrivateKey)
	}

	// Open Database connection.
	db, err := database.Open(database.DatabaseLocation)
	if err != nil {
		return err
	}

	go database.GetAllOld(db, output)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Remove and notify all nodes who haven't checked in.
	for node := range output {
		err = database.Remove(tx, node.ID)
		if err != nil {
			return err
		}

		if node.Email != "" && config.Mail.Setup {
			content, err := ioutil.ReadFile("templates/email/alert_message.txt")
			if err != nil {
				log.Fatal(err)
			}
			emailMessage := string(content)

			message := mg.NewMessage(
				config.Mail.Sender,
				"Arken Node Offline",
				emailMessage,
				config.CleanEmail(node.Email))

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			_, _, err = mg.Send(ctx, message)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
