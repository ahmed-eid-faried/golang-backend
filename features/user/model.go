package user

import (
	"log"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"

	sqldb "main/core/db/sql"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// sqldb "golang-backend/core/db/sqldb"

// User represents a user model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// InitDataBase initializes the database tables and generates fake and provided users.
func InitData() {
	// Create the users table if it doesn't exist
	sqldb.CreateTable("users", "id SERIAL PRIMARY KEY, username VARCHAR(255), email VARCHAR(255)")

	// Generate 20 fake users
	GenerateFakeUsers(20)

	// Define the list of users
	users := []User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
		// Add more users as needed
	}

	// Generate the provided users
	GenerateUsers(users)
}

func GenerateFakeUsers(numUsers int) {
	for i := 0; i < numUsers; i++ {

		var user User
		err := faker.FakeData(&user)
		if err != nil {
			log.Println("Error generating fake data for user:", err)
			continue
		}
		_, err = sqldb.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
		if err != nil {
			log.Println("Error inserting fake user data:", err)
			continue
		}

	}
}

func GenerateUsers(users []User) {
	for _, user := range users {
		_, err := sqldb.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
		if err != nil {
			log.Println("Error inserting user data:", err)
			continue
		}
	}
}
