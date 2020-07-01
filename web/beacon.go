package web

import (
	"encoding/json"
	"net/http"

	"github.com/arkenproject/stats/database"
)

// Beacon is the api front end for nodes to check in with the stats server.
func Beacon(w http.ResponseWriter, r *http.Request) {
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
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
