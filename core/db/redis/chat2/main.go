package chat2

import "github.com/gin-gonic/gin"

func InitChat(r *gin.Engine, v1 *gin.RouterGroup) {

	// Load HTML page
	r.LoadHTMLFiles("./core/db/redis/chat2/index.html")

	// Serve index.html
	v1.GET("/chat2", HandleIndex)

	// Serve index.html
	v1.GET("/chat2/path", HandleChatPath)

	// Handle WebSocket connections
	v1.GET("/ws", HandleWebSocket)

	// Start broadcasting messages
	go Listen()
}
