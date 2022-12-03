package controllers

import (
	"chatapp/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.

var (
	Rooms = map[int64]*Hub{}
)

type Hub struct {
	// Registered clients.
	clients []*Client
	chat    *models.Chat

	// Inbound messages from the clients.
	broadcast chan *BroadcastEvent

	// Register/Unregister requests from the clients.
	register   chan *Client
	unregister chan *Client
}

func newHub(chat *models.Chat) *Hub {
	return &Hub{
		broadcast:  make(chan *BroadcastEvent),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    []*Client{},
		chat:       chat,
	}
}

func (h *Hub) run() {
	for loop := true; loop; {
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
			if len(h.clients) == 0 {
				delete(Rooms, h.chat.Id)
				close(h.register)
				close(h.unregister)
				close(h.broadcast)
				loop = false
			}
		case event := <-h.broadcast:
			for _, client := range h.clients {
				client.send <- event
			}
		}
	}
}
