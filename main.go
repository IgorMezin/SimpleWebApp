package main

import (
	//"firstWebApp/database"
	"firstWebApp/registration"
	"log"
	"net/http"
)

func main() {
	//database.InitDB()
	//defer database.DB.Close()

	http.HandleFunc("/", registration.Handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
