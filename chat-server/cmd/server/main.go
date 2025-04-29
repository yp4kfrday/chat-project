package main

import (
	"log"
	"net/http"

	"chat-server/internal/server"
)

func main() {
	log.Println("[INFO] Starting chat server on :8080...")

	hub := server.NewServer()

	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWs(hub, w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("[FATAL] Server failed:", err)
	}
}