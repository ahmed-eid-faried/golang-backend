
// models/room.go
package chat

type Room struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Members []string `json:"members"`
    // Add more fields as needed
}
