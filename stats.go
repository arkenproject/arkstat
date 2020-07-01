package main

import (
	"fmt"
	"net/http"

	"github.com/arkenproject/stats/tasks"

	"github.com/arkenproject/stats/web"
)

func main() {

	go tasks.Start()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("templates/fonts"))))

	http.HandleFunc("/beacon", web.Beacon)
	http.HandleFunc("/", web.Page)

	fmt.Println("Started Stats Webserver!")
	http.ListenAndServe(":8080", nil)
}
