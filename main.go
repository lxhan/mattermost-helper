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
	router.GET("/daily-pt", DailyPT)
	router.GET("/reminder/:type", Reminder)
	router.GET("/reminder-pt/:type", ReminderPT)

	log.Fatal(http.ListenAndServe(":8080", router))
}
