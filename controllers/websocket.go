package controllers

import (
	"bytes"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"

	"chatapp/models"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	BaseController
}

// Join method handles WebSocket requests for WebSocketController.
func (c *WebSocketController) Join() {
	chatid, err := c.GetInt64("chat")
	fmt.Println("go", err, chatid, c.Userinfo)

	if err != nil || c.Userinfo == nil {
		fmt.Println("No user or chat")
		c.Data["json"] = bson.M{"Fuck": "you", "why": "you"}
		c.ServeJSON()
		return
	}

	chat, err := models.IsAllowed(chatid, c.Userinfo.Id)
	if err != nil {
		fmt.Println("Not Allowed")
		c.Data["json"] = bson.M{"Fuck": "you"}
		c.ServeJSON()
		return
	}
	fmt.Println("Allowed")

	// Upgrade from http request to WebSocket.
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Println("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	client := Join(chat, c.Userinfo, ws)
	defer Leave(client)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		client.hub.broadcast <- message
	}
}
