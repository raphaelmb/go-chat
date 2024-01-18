package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelmb/go-chat/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
