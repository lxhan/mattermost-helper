package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := httprouter.New()

	router.GET("/ping", Ping)
	router.GET("/daily", Daily)
	router.GET("/reminder/:type", Reminder)

	log.Fatal(http.ListenAndServe(":8080", router))
}
