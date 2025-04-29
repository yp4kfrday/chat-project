package server

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-server/internal/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	Server *Server
	Conn   *websocket.Conn
	Send   chan []byte
	Username string
}

func ServeWs(server *Server, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[ERROR] WebSocket upgrade:", err)
		return
	}

	client := &Client{
		Server: server,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Username: username,
	}

	server.Register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
    defer func() {
        c.Server.Unregister <- c
        c.Conn.Close()
    }()

    for {
        _, data, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println("[ERROR] Read error:", err)
            break
        }

        var message models.Message
        err = json.Unmarshal(data, &message)
        if err != nil {
            log.Println("[ERROR] JSON Unmarshal error:", err)
            continue
        }

        msgBytes, _ := json.Marshal(message)
        c.Server.Broadcast <- msgBytes
    }
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("[ERROR] Write:", err)
			break
		}
	}
}