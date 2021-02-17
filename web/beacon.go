package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/arken/arkstat/config"
	"github.com/arken/arkstat/database"
	"github.com/mailgun/mailgun-go/v4"
)

// Beacon is the api front end for nodes to check in with the arkstat server.
func Beacon(w http.ResponseWriter, r *http.Request) {
	new := false

	if r.Method == http.MethodPost {
		var node database.Node
		err := json.NewDecoder(r.Body).Decode(&node)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db, err := database.Open(database.DatabaseLocation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = database.Get(db, node.ID)
		if err != nil {
			new = true
		}

		err = database.Add(db, node)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		out, err := database.Get(db, node.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(out)

		if new && out.Email != "" && config.Mail.Setup {
			mg := mailgun.NewMailgun(config.Mail.Domain, config.Mail.PrivateKey)
			content, err := ioutil.ReadFile("templates/email/hello_message.txt")
			if err != nil {
				log.Fatal(err)
			}
			emailMessage := string(content)

			message := mg.NewMessage(
				config.Mail.Sender,
				"Welcome to Arken",
				emailMessage,
				config.CleanEmail(node.Email))

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			_, _, err = mg.Send(ctx, message)
			if err != nil {
				log.Println(err)
			}
		}
		err = db.Close()
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
