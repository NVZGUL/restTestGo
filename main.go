package main

import (
	"log"
	"net/http"

	"github.com/NVZGUL/restTestGo/handlers"
)

func main() {
	courses := handlers.NewCourses()
	http.DefaultServeMux.HandleFunc("/", courses.Handle)
	err = http.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
