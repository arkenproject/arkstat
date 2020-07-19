package main

import (
	"fmt"
	"net/http"

	"github.com/arkenproject/arkstat/tasks"
	"github.com/arkenproject/arkstat/web"
)

func main() {

	go tasks.Start()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("templates/fonts"))))
	http.Handle("/graphic/", http.StripPrefix("/graphic/", http.FileServer(http.Dir("templates/graphic"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("templates/images"))))

	http.HandleFunc("/beacon", web.Beacon)
	http.HandleFunc("/", web.Page)

	fmt.Println("Started ArkStat Webserver!")
	http.ListenAndServe(":8080", nil)
}
