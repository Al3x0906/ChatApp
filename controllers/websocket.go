package controllers

import (
	"log"
	"net/http"
	"regexp"

	"chatapp/models"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	BaseController
}

// Join method handles WebSocket requests for WebSocketController.
func (c *WebSocketController) Join() {
	defer c.ServeJSON()

	chatid, err := c.GetInt64("chat")

	if err != nil || c.Userinfo == nil {
		log.Println("No user or chat")
		return
	}

	chat, err := models.IsAllowed(chatid, c.Userinfo.Id)
	if err != nil {
		log.Println("Not allowed", chatid, c.Userinfo.Id)
		return
	}
	log.Println("welcome")

	// Upgrade from http request to WebSocket.
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			re := regexp.MustCompile(`^http://(localhost|127\.0\.0\.1):(3000|8080)`)
			matches := re.FindStringSubmatch(r.Header.Get("Origin"))
			return len(matches) > 0
		},
	}

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	// Join chat room.
	client := Join(chat, c.Userinfo, ws)
	go client.writePump()
	go client.readPump()

}
