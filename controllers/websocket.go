package controllers

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"chatapp/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	BaseController
}

// Join method handles WebSocket requests for WebSocketController.
func (c *WebSocketController) Join() {
	chat, _ := c.GetInt64("chat")
	if !models.IsAllowed(c.Userinfo.Id, chat) {
		fmt.Println("Not Allowed", chat, c.Userinfo.Id)
		c.Data["json"] = bson.M{"Fuck": "you"}
		c.ServeJSON()
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Println("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	Join(chat, ws)
	defer Leave(chat)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EventMessage, chat, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
