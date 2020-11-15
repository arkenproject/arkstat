package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arkenproject/arkstat/tasks"
	"github.com/arkenproject/arkstat/web"
)

// HTMLDir wraps a http.Dir stuct to add custom opening formats.
type HTMLDir struct {
	d http.Dir
}

func main() {

	go tasks.Start()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("templates/fonts"))))
	http.Handle("/graphic/", http.StripPrefix("/graphic/", http.FileServer(http.Dir("templates/graphic"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("templates/images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js"))))
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(HTMLDir{http.Dir("templates/docs")})))
	http.Handle("/favicon.ico", http.FileServer(http.Dir("templates/images")))
	http.Handle("/robots.txt", http.FileServer(http.Dir("templates/static")))
	http.Handle("/sitemap.xml", http.FileServer(http.Dir("templates/static")))

	http.HandleFunc("/beacon", web.Beacon)
	http.HandleFunc("/", web.Page)

	fmt.Println("Started ArkStat Webserver!")
	http.ListenAndServe(":8080", nil)
}

// Open defines how to open specified http files.
func (d HTMLDir) Open(name string) (http.File, error) {
	// Try name with added extension
	f, err := d.d.Open(name + ".html")
	if os.IsNotExist(err) {

		// Not found, try again with name as supplied.
		if f, err := d.d.Open(name); err == nil {
			return f, nil
		}
	}

	return f, err
}
