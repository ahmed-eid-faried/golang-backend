// main.go
package chat2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// HandleWebSocket handles WebSocket requests
func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var message Message
		err := ws.ReadJSON(&message)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- message
	}
}

// HandleIndex renders the index.html file
// @Summary Get chat for chat realtime
// @Produce html
// @Success 200 {string} string "Path for chat realtime"
// @Router /chat2/ [get]
// @Tags Chat RealTime
func HandleIndex(c *gin.Context) {
	c.File("./core/db/redis/chat2/index.html")
}

// HandleChatPath handles the endpoint to get the path for chat realtime
// @Summary Get Path for chat realtime
// @Produce json
// @Success 200 {object} string "Path for chat realtime"
// @Router /chat2/path [get]
// @Tags Chat RealTime
func HandleChatPath(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"path": "http://localhost:8081/api/v1/chat2"})
}

// Listen broadcasts messages to all connected clients
func Listen() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
