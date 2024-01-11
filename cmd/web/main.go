package main

import (
	"log"
	"net/http"

	"github.com/raphaelmb/go-chat/internal/handlers"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	log.Println("Starting channel listener")
	go handlers.ListToWsChannel()

	log.Println("Started server on port :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
