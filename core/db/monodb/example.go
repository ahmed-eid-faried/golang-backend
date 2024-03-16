package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func example() {
	// Initialize the MongoDB connection
	InitMongoDB()

	// Sample data
	user := bson.M{"name": "John Doe", "age": 30}
	users := []interface{}{
		bson.M{"name": "Alice", "age": 25},
		bson.M{"name": "Bob", "age": 35},
	}

	address := bson.M{"street": "123 Main St", "city": "New York"}

	// CRUD operations on the 'users' collection
	err := Create("users", user)
	if err != nil {
		log.Fatal(err)
	}

	err = Add("users", users)
	if err != nil {
		log.Fatal(err)
	}

	// Perform search
	var searchQuery interface{}
	searchQuery = bson.M{"search": bson.M{"query": "John"}}
	results, err := Search("users", searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Search Results:", results)

	// Retrieve a user
	filter := bson.M{"name": "Alice"}
	result, err := View("users", filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("View Result:", result)

	// Update a user
	update := bson.M{"$set": bson.M{"age": 26}}
	err = Update("users", filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Delete users
	deleteFilter := bson.M{"age": bson.M{"$gte": 35}} // Delete users aged 35 and above
	err = Delete("users", deleteFilter)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve all users
	allUsers, err := ViewAll("users")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All Users:", allUsers)

	// Perform aggregation operations
	// Example: Group by age
	groupStage := bson.M{"_id": "$age", "count": bson.M{"$sum": 1}}
	groupResults, err := GroupBy("users", "age")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Group Results:", groupResults, groupStage)

	// Example: Order by name in descending order
	orderByResults, err := OrderBy("users", "name", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ordered Results:", orderByResults)

	// CRUD operations on the 'address' collection
	err = Create("address", address)
	if err != nil {
		log.Fatal(err)
	}

	// Perform schema validation for the 'address' collection
	validationRules := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"street", "city"},
		},
	}
	err = SetSchemaValidation("address", validationRules)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Schema validation set for 'address' collection.")
}

// Your CRUD and aggregation functions...
