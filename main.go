package main

import (
	"firstWebApp/registration"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", registration.Handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
