// main.go
package chat

import (
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"

	kredis "main/core/db/redis"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// redis "main/db/redis"
// var redisClient *redis.Client

// func ChatRoute(r *gin.Engine) {
// 	kredis.InitDatabase()

// 	// Serve static files
// 	r.Static("/static", "./static")

// 	// Load HTML page
// 	r.LoadHTMLFiles("./views/index.html")

// 	// Websocket endpoint
// 	r.GET("/ws", func(c *gin.Context) {
// 		handleWebsocket(c.Writer, c.Request)
// 	})

// 	// // Start Gin server
// 	// if err := r.Run(":8080"); err != nil {
// 	// 	log.Fatal("Failed to start server: ", err)
// 	// }
// }

// @Summary Handle websocket connection
// @Description Upgrade HTTP connection to websocket and handle incoming messages
// @Success 200 {string} string "OK"
// @Router /ws [get]
func HandleWebsocket(c *gin.Context) {
	// Upgrade HTTP connection to websocket
	conn, err := kredis.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	// Handle incoming messages
	for {
		// Read message from websocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Broadcast message to all clients
		err = kredis.RedisClient.Publish(kredis.CTX, "chat", msg).Err()
		if err != nil {
			log.Println("Error publishing message to Redis:", err)
			break
		}
	}
}
