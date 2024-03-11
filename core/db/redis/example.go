package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Example() {
	// Initialize the database
	if err := InitDatabase(); err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	// Cache a value
	if err := CacheValue("key3", "value3", 10*time.Second); err != nil {
		fmt.Println("Error caching value:", err)
		return
	}

	// Remove a cached value
	if err := RemoveCachedValue("key1"); err != nil {
		fmt.Println("Error removing cached value:", err)
		return
	}

	// Example: Retrieve a value from Redis
	val, err := RedisClient.Get(CTX, "key1").Result()
	if err == redis.Nil {
		fmt.Println("Key1 does not exist")
	} else if err != nil {
		fmt.Println("Error retrieving value for key1:", err)
	} else {
		fmt.Println("Value for key1:", val)
	}
	////////////////////////////////  chat  \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	// Example usage
	roomID := "chatroom1"
	userID := 123

	// Join user to room
	if err := JoinRoom(roomID, userID); err != nil {
		fmt.Println("Error joining room:", err)
		return
	}

	fmt.Println("User", userID, "joined room", roomID)

	// Send a message
	message := "Hello, World!"
	if err := SendMessage(roomID, userID, message); err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("User", userID, "sent message:", message)

	// Fetch messages in the room
	messages, err := GetRoomMessages(roomID)
	if err != nil {
		fmt.Println("Error fetching messages in room:", err)
		return
	}

	fmt.Println("Messages in room", roomID, ":", messages)

	// Leave user from room
	if err := LeaveRoom(roomID, userID); err != nil {
		fmt.Println("Error leaving room:", err)
		return
	}

	fmt.Println("User", userID, "left room", roomID)
}
