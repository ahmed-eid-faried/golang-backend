package address

import (
	"log"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"

	sqldb "main/core/db/sql"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// sqldb "golang-backend/core/db/sqldb"

// // Address represents an address model
// type Address struct {
// 	ID         int    `json:"id"`
// 	Street     string `json:"street"`
// 	City       string `json:"city"`
// 	State      string `json:"state"`
// 	PostalCode string `json:"postal_code"`
// 	UserID     int    `json:"user_id"`
// }

// Address represents the structure of an address
type Address struct {
	ID         int    `json:"id" example:"1" format:"int32"`
	Street     string `json:"street" example:"123 Example St."`
	City       string `json:"city" example:"Example City"`
	State      string `json:"state" example:"Example State"`
	PostalCode string `json:"postal_code" example:"12345"`
	UserID     int    `json:"user_id" example:"1" format:"int32"`
}

// InitDataBase initializes the database tables and generates fake and provided addresses.
func InitData() {
	// Create the addresses table if it doesn't exist
	sqldb.CreateTable("address", "id SERIAL PRIMARY KEY, street VARCHAR(255), city VARCHAR(255), state VARCHAR(255), postal_code VARCHAR(255), user_id INT")

	// Define the list of addresses
	addresses := []Address{
		{
			ID:         1, //don't use this ID in your code
			Street:     "Street",
			City:       "City",
			State:      "State",
			PostalCode: "PostalCode",
			UserID:     1,
		},
		{
			ID:         2, //don't use this ID in your code
			Street:     "Street",
			City:       "City",
			State:      "State",
			PostalCode: "PostalCode",
			UserID:     2,
		},
	}

	// Generate the provided users
	GenerateAddresses(addresses)

	// Generate 20 fake users
	GenerateFakeAddresses(20, 3)
}

// GenerateFakeAddresses generates fake addresses
func GenerateFakeAddresses(numAddresses int, userID int) []Address {
	var addresses []Address
	for i := 0; i < numAddresses; i++ {
		var address Address
		err := faker.FakeData(&address)
		if err != nil {
			log.Println("Error generating fake data for address:", err)
			continue
		}
		address.UserID = userID
		addresses = append(addresses, address)
	}
	return addresses
}

// AddAddresses adds a list of addresses to the database
func GenerateAddresses(addresses []Address) {
	for _, address := range addresses {
		// Insert address into the database
		_, err := sqldb.DB.Exec("INSERT INTO address (street, city, state, postal_code, user_id) VALUES ($1, $2, $3, $4, $5)", address.Street, address.City, address.State, address.PostalCode, address.UserID)
		if err != nil {
			log.Println("Error inserting address data:", err)
			continue
		}
	}
}
