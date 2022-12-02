package controllers

import "chatapp/models"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.

var (
	Rooms map[int64]*Hub
)

type Hub struct {
	// Registered clients.
	clients []*Client
	chat    *models.Chat

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register/Unregister requests from the clients.
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    []*Client{},
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if h.chat.User1 == client.user.Id || h.chat.User2 == client.user.Id {
				h.clients = append(h.clients, client)
			}
		case client := <-h.unregister:
			for i, cl := range h.clients {
				if cl == client {
					h.clients[i] = h.clients[len(h.clients)-1]
					h.clients = h.clients[:len(h.clients)-1]
					break
				}
			}
		case message := <-h.broadcast:
			for _, client := range h.clients {
				client.send <- message
			}
		}
	}
}
