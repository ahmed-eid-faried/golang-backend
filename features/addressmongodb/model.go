package addressmongodb

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"

	mongodb "main/core/db/monodb"
)

// Address represents the structure of an address
type Address struct {
	// ID         int    `json:"id" example:"1" format:"int32"`
	ID         string `json:"id" bson:"_id,omitempty"`
	Street     string `json:"street" example:"123 Example St."`
	City       string `json:"city" example:"Example City"`
	State      string `json:"state" example:"Example State"`
	PostalCode string `json:"postal_code" example:"12345"`
	UserID     string `json:"user_id" example:"1"`
	// UserID     string `json:"user_id" example:"1" format:"int32"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided addresses.
func InitData() {
	mongodb.KInit()
	// var DBColl *mongo.Collection

	// // Initialize MongoDB client and connect
	// client, err := mongodb.NewClient()
	// if err != nil {
	// 	log.Fatal("Error initializing MongoDB client:", err)
	// }
	// defer client.Disconnect(context.Background())

	// // Get a handle to the database
	// db := client.Database("your_database_name")

	// // Create the addresses collection if it doesn't exist
	// collection := db.Collection("address")

	// // Define the list of addresses
	addresses := []interface{}{
		Address{
			Street:     "Street",
			City:       "City",
			State:      "State",
			PostalCode: "PostalCode",
			UserID:     "1",
		},
		Address{
			Street:     "Street",
			City:       "City",
			State:      "State",
			PostalCode: "PostalCode",
			UserID:     "2",
		},
	}

	// Insert provided addresses into MongoDB
	_, err := mongodb.DB.Collection("address").InsertMany(CTX, addresses)

	if err != nil {
		log.Println("Error inserting fake addresses:", err)
	}

	// Generate and insert fake addresses
	fakeAddresses := GenerateFakeAddresses(20)
	_, err = mongodb.DB.Collection("address").InsertMany(CTX, fakeAddresses)
	if err != nil {
		log.Println("Error inserting fake addresses:", err)
	}
}

// GenerateFakeAddresses generates fake addresses
func GenerateFakeAddresses(numAddresses int) []interface{} {
	var addresses []interface{}
	for i := 0; i < numAddresses; i++ {
		var address Address
		err := faker.FakeData(&address)
		if err != nil {
			log.Println("Error generating fake data for address:", err)
			continue
		}
		address.UserID = fmt.Sprintf("%s", rand.Int())
		addresses = append(addresses, address)
	}
	return addresses
}

// // GetAllAddresses retrieves all addresses from MongoDB
// func GetAllAddresses() ([]Address, error) {
// 	// Initialize MongoDB client and connect
// 	client, err := mongodb.NewClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	// Get a handle to the database
// 	db := client.Database("your_database_name")

// 	// Get a handle to the addresses collection
// 	collection := db.Collection("address")

// 	// Find all addresses
// 	cursor, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.Background())

// 	// Decode addresses from cursor
// 	var addresses []Address
// 	if err := cursor.All(context.Background(), &addresses); err != nil {
// 		return nil, err
// 	}

// 	return addresses, nil
// }
