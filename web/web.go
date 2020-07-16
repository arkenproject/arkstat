package web

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

// Values is a structure of variables to be replaced in the template.
type values struct {
	Title       string
	UsedSpace   string
	TotalSpace  string
	ActiveNodes int
}

// PageValues is a stucture of data to be pasted into web page template.
var PageValues values

func init() {
	PageValues.Title = "Arken"

}

// Page is the default web handler to render the "arkstat" html page.
func Page(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}
	// Return a 404 if the request is for a directory
	if info.IsDir() && fp != "templates" {
		http.NotFound(w, r)
		return
	}

	t = template.New("index.html")

	t, _ = t.ParseFiles("templates/index.html")

	t.Execute(w, PageValues)

}
