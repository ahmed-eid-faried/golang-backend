package redis

import (
	// "context"
	"fmt"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	// "github.com/go-redis/redis/v8"
)
var (
	// users    = make(map[string]*User)
	// messages []Message
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)
// Mutex for synchronizing room operations
var roomMutex sync.Mutex

// JoinRoom adds a user to a specified room
func JoinRoom(roomID string, userID int) error {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	// Add user to room set
	return RedisClient.SAdd(CTX, "room:"+roomID, userID).Err()
}

// LeaveRoom removes a user from a specified room
func LeaveRoom(roomID string, userID int) error {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	// Remove user from room set
	return RedisClient.SRem(CTX, "room:"+roomID, userID).Err()
}

// GetUsersInRoom retrieves the list of users in a specified room
func GetUsersInRoom(roomID string) ([]int, error) {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	// Get all users in the room
	members, err := RedisClient.SMembers(CTX, "room:"+roomID).Result()
	if err != nil {
		return nil, err
	}

	// Convert member strings to integers
	var users []int
	for _, member := range members {
		userID, _ := strconv.Atoi(member)
		users = append(users, userID)
	}

	return users, nil
}

// SendMessage sends a message to a specified room
func SendMessage(roomID string, userID int, message string) error {
	// Push message to the end of the room's message list
	return RedisClient.RPush(CTX, "messages:"+roomID, fmt.Sprintf("%d:%s", userID, message)).Err()
}

// GetRoomMessages retrieves the messages in a specified room
func GetRoomMessages(roomID string) ([]string, error) {
	// Get all messages in the room
	return RedisClient.LRange(CTX, "messages:"+roomID, 0, -1).Result()
}
