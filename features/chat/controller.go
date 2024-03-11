// // controller/chat_controller.go

package chat

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"
//
//
//
//
//

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8" // Import Redis client library
// 	"github.com/gorilla/websocket"
// )

// func NewChatController() *ChatController {
// 	return &ChatController{}
// }

// // RedisClient holds the Redis client instance
// var RedisClient *redis.Client

// // InitializeRedis initializes the Redis client
// func InitializeRedis() {
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Redis server address
// 		Password: "",               // Redis password
// 		DB:       0,                // Default database
// 	})
// }

// // Set a context with timeout, cancelation, or background
// var ctx = context.Background() // This creates a background context
// var cancel context.CancelFunc

// // Save message to the database (Redis)
// func saveMessageToRedis(roomId string, message Message) error {
// 	// Convert message to JSON
// 	msgBytes, err := json.Marshal(message)
// 	if err != nil {
// 		return err
// 	}
// 	// Set a context with a timeout of 5 seconds
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel() // Always remember to defer the cancelation to prevent leaks

// 	// Save message to Redis list
// 	_, err = RedisClient.LPush(ctx, "room:"+roomId, msgBytes).Result()
// 	return err
// }

// // Fetch messages from the database based on roomId (Redis)
// func fetchMessagesFromRedis(roomId string) ([]Message, error) {
// 	// Set a context with a timeout of 5 seconds
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel() // Always remember to defer the cancelation to prevent leaks

// 	// Fetch messages from Redis list
// 	msgBytes, err := RedisClient.LRange(ctx, "room:"+roomId, 0, -1).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Unmarshal messages
// 	var messages []Message
// 	for _, msg := range msgBytes {
// 		var message Message
// 		err := json.Unmarshal([]byte(msg), &message)
// 		if err != nil {
// 			log.Println("Error unmarshalling message:", err)
// 			continue
// 		}
// 		messages = append(messages, message)
// 	}

// 	return messages, nil
// }

// // ChatController handles chat-related operations
// type ChatController struct{}

// // GetChatRoom gets messages from a chat room

// // @Summary Get messages from a chat room
// // @Description Get all messages from a specific chat room
// // @Tags chat
// // @Param roomId path string true "Chat Room ID"
// // @Success 200 {object} []Message
// // @Failure 404 {object} ErrorResponse
// // @Router /chat/{roomId} [get]
// func (c *ChatController) GetChatRoom(ctx *gin.Context) {
// 	roomId := ctx.Param("roomId")
// 	messages, err := fetchMessagesFromRedis(roomId)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, messages)
// }

// // SendMessage sends a message to a chat room

// // @Summary Send a message to a chat room
// // @Description Send a message to a specific chat room
// // @Tags chat
// // @Param roomId path string true "Chat Room ID"
// // @Param message body Message true "Message Object"
// // @Success 200 {object} SuccessResponse
// // @Failure 400 {object} ErrorResponse
// // @Router /chat/{roomId} [post]
// func (c *ChatController) SendMessage(ctx *gin.Context) {
// 	roomId := ctx.Param("roomId")
// 	var message Message
// 	if err := ctx.BindJSON(&message); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := saveMessageToRedis(roomId, message); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"status": "Message sent successfully"})
// }

// type Message struct {
// 	Text string `json:"text"`
// 	User string `json:"user"`
// }

// type SuccessResponse struct {
// 	Status string `json:"status"`
// }

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

// // User struct represents a user in the chat system
// type User struct {
// 	ID       string `json:"id"`
// 	Username string `json:"username"`
// 	Conn     *websocket.Conn
// }

// var (
// 	users    = make(map[string]*User)
// 	messages []Message
// 	upgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}
// )

// // WebSocketHandler upgrades HTTP connections to WebSocket and handles WebSocket events
// // @Summary Upgrade HTTP connections to WebSocket
// // @Description This endpoint upgrades HTTP connections to WebSocket and handles WebSocket events
// // @Tags chat
// // @Router /chat/ws [get]

// func WebSocketHandler(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println("Failed to upgrade to WebSocket:", err)
// 		return
// 	}

// 	defer conn.Close()

// 	user := &User{
// 		ID:   "unique_user_id",
// 		Conn: conn,
// 	}

// 	users[user.ID] = user

// 	for {
// 		var msg Message
// 		_, msgBytes, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Error reading message:", err)
// 			break
// 		}
// 		if err := json.Unmarshal(msgBytes, &msg); err != nil {
// 			log.Println("Error unmarshalling message:", err)
// 			continue
// 		}
// 		broadcastMessage(msg)
// 	}
// }

// // Broadcasts message to all connected WebSocket clients
// func broadcastMessage(msg Message) {
// 	// Add message to the message history
// 	messages = append(messages, msg)
// 	// Convert message to JSON
// 	msgBytes, err := json.Marshal(msg)
// 	if err != nil {
// 		log.Println("Error marshalling message:", err)
// 		return
// 	}
// 	// Send message to all connected clients
// 	for _, user := range users {
// 		err := user.Conn.WriteMessage(websocket.TextMessage, msgBytes)
// 		if err != nil {
// 			log.Println("Error sending message to user:", err)
// 			continue
// 		}
// 	}
// }
