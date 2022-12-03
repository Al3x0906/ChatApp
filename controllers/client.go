package controllers

import (
	"chatapp/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

type BroadcastEvent struct {
	Type_   models.EventType `json:"type"`
	Message *models.Message  `json:"message"`
}

func Join(chat *models.Chat, user *models.User, ws *websocket.Conn) *Client {
	hub, ok := Rooms[chat.Id]
	if !ok {
		hub = newHub(chat)
		Rooms[chat.Id] = hub
		go hub.run()
	}
	client := &Client{hub: hub, user: user, conn: ws, send: make(chan *BroadcastEvent)}
	hub.register <- client
	return client
}

func Leave(client *Client) {
	hub := client.hub
	hub.unregister <- client
	close(client.send)
	_ = client.conn.Close()
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	user *models.User
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan *BroadcastEvent
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		Leave(c)
	}()

	for {
		_, req, err := c.conn.ReadMessage()
		var m bson.M
		errUnmarshal := json.Unmarshal(req, &m)

		if err != nil {
			break
		}
		if errUnmarshal != nil {
			continue
		}

		t, ok := m["type"].(string)
		if !ok {
			continue
		}
		switch t {
		case models.EventMessage:
			content, ok := m["content"].(string)
			if !ok {
				continue
			}
			if content := strings.TrimSpace(content); content == "" {
				continue
			}
			receiver := models.IfThenElse(c.user.Id == c.hub.chat.User1, c.hub.chat.User2, c.hub.chat.User1)
			mess := &models.Message{Chat: c.hub.chat.Id, Sender: c.user.Id, Receiver: receiver.(int64), Content: []byte(content), Sent: time.Now()}
			err = mess.Insert()
			if err != nil {
				continue
			}
			c.hub.broadcast <- &BroadcastEvent{models.EventMessage, mess}

		case models.EventDelete:
			id, ok := m["id"].(float64)
			if !ok {
				continue
			}

			mess := &models.Message{Id: int64(id), Chat: c.hub.chat.Id, Sender: c.user.Id}
			count, err := mess.Delete()
			if err != nil || count == 0 {
				continue
			}
			c.hub.broadcast <- &BroadcastEvent{models.EventDelete, mess}

		case models.EventEdit:
			id, ok := m["id"].(float64)
			if !ok {
				continue
			}
			content, ok := m["content"].(string)
			if !ok {
				continue
			}
			if content := strings.TrimSpace(content); content == "" {
				continue
			}

			mess := &models.Message{Id: int64(id), Chat: c.hub.chat.Id, Sender: c.user.Id, Content: []byte(content)}
			count, err := mess.Update()
			if err != nil || count == 0 {
				continue
			}
			c.hub.broadcast <- &BroadcastEvent{models.EventEdit, mess}
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	defer func() {
		_ = c.conn.Close()
	}()
	for {
		select {
		case event, ok := <-c.send:

			// if socket closed
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteJSON(event)
			if err != nil {
				return
			}

			if event.Type_ == models.EventMessage && event.Message.Receiver != c.user.Id && event.Message.Seen == models.ZeroTime {
				event.Message.Seen = time.Now()
				err := event.Message.SetSeen()
				if err != nil {
					continue
				}
				c.hub.broadcast <- &BroadcastEvent{models.EventSeen, event.Message}
			}
		}
	}
}
