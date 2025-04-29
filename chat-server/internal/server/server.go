package server

import (
	"chat-server/internal/models"
	"encoding/json"
)

type Server struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.Register:
			s.Clients[client] = true
			welcome := models.Message{Username: "Server", Content: "Welcome " + client.Username + "!"}
			msgBytes, _ := json.Marshal(welcome)
			client.Send <- msgBytes

		case client := <-s.Unregister:
			if _, ok := s.Clients[client]; ok {
				delete(s.Clients, client)
				close(client.Send)
			}

		case message := <-s.Broadcast:
			for client := range s.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(s.Clients, client)
				}
			}
		}
	}
}