package main

import (
	"PlayerElo/handlers"

	muxHandlers "github.com/gorilla/handlers"

	"log"
	"net/http"
	"os"
)

func main() {
	s := handlers.New(handlers.SiteProperties{
		Title:    "EloMapper",
		Subtitle: "Tracking stats",
	})
	http.Handle("/", muxHandlers.LoggingHandler(os.Stdout, s.Router()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
