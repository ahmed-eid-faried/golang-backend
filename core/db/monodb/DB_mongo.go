package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Init initializes the Mongodb connection

var (
	Client *mongo.Client
	DB     *mongo.Database
	// DBColl *mongo.Collection
	MonHelper *MongoDBHelper
)
//  dbName :="goblog"

// Init initializes the MongoDB connection
func KInit() {
	// Set Client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	DB = client.Database("goblog") // Specify the database name

	log.Println("Connected to MongoDB")
}
func Init() {
	// Set Client options
	ClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(context.Background(), ClientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	DB = Client.Database("goblog") // Specify the database name
}

// Create inserts a new document into the specified collection
func Create(collectionName string, document interface{}) error {
	collection := DB.Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), document)
	return err
}

// Add inserts multiple documents into the specified collection
func Add(collectionName string, documents []interface{}) error {
	collection := DB.Collection(collectionName)
	_, err := collection.InsertMany(context.Background(), documents)
	return err
}

// Delete deletes documents from the specified collection based on the filter
func Delete(collectionName string, filter bson.M) error {
	collection := DB.Collection(collectionName)
	_, err := collection.DeleteMany(context.Background(), filter)
	return err
}

// Update updates documents in the specified collection based on the filter
func Update(collectionName string, filter bson.M, update bson.M) error {
	collection := DB.Collection(collectionName)
	_, err := collection.UpdateMany(context.Background(), filter, update)
	return err
}

// Max finds the maximum value for the specified field in the collection
func Max(collectionName, fieldName string) (interface{}, error) {
	collection := DB.Collection(collectionName)
	opts := options.FindOne().SetSort(bson.D{{Key: fieldName, Value: -1}})
	var result bson.M
	if err := collection.FindOne(context.Background(), bson.D{}, opts).Decode(&result); err != nil {
		return nil, err
	}
	return result[fieldName], nil
}

// Min finds the minimum value for the specified field in the collection
func Min(collectionName, fieldName string) (interface{}, error) {
	collection := DB.Collection(collectionName)
	opts := options.FindOne().SetSort(bson.D{{Key: fieldName, Value: 1}})
	var result bson.M
	if err := collection.FindOne(context.Background(), bson.D{}, opts).Decode(&result); err != nil {
		return nil, err
	}
	return result[fieldName], nil
}

// GroupBy groups documents in the collection based on the specified field
func GroupBy(collectionName string, field string) ([]bson.M, error) {
	collection := DB.Collection(collectionName)
	pipeline := []bson.M{
		{"$group": bson.M{"_id": "$" + field, "count": bson.M{"$sum": 1}}},
	}
	cur, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []bson.M
	if err := cur.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

// OrderBy retrieves documents from the specified collection ordered by the specified field
func OrderBy(collectionName, fieldName string, desc bool) ([]bson.M, error) {
	collection := DB.Collection(collectionName)
	order := 1
	if desc {
		order = -1
	}
	opts := options.Find().SetSort(bson.D{{Key: fieldName, Value: order}})
	cur, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []bson.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func Group(databaseName, collectionName string, groupStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$group", Value: groupStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Limit performs $limit aggregation operation
func Limit(databaseName, collectionName string, limit int64) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$limit", Value: limit}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Project performs $project aggregation operation
func Project(databaseName, collectionName string, projectStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$project", Value: projectStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Sort performs $sort aggregation operation
func Sort(databaseName, collectionName string, sortStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$sort", Value: sortStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Match performs $match aggregation operation
func Match(databaseName, collectionName string, matchStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$match", Value: matchStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// AddFields performs $addFields aggregation operation
func AddFields(databaseName, collectionName string, addFieldsStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$addFields", Value: addFieldsStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Count performs $count aggregation operation
func Count(databaseName, collectionName string) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$count", Value: "total"}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Lookup performs $lookup aggregation operation
func Lookup(databaseName, collectionName string, lookupStage interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$lookup", Value: lookupStage}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Out performs $out aggregation operation
func Out(databaseName, collectionName, outputCollection string) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$out", Value: outputCollection}}}
	return runAggregation(databaseName, collectionName, pipeline)
}

// Search utilizes the $search operator in an aggregation pipeline
func Search(collectionName string, searchQuery interface{}) ([]bson.M, error) {
	pipeline := mongo.Pipeline{{{Key: "$search", Value: searchQuery}}}
	return runAggregation(kdatabaseName, collectionName, pipeline)
}

// // Search retrieves documents from the specified collection based on the filter
// func Search(collectionName string, filter bson.M) ([]bson.M, error) {
//     collection := DB.Collection(collectionName)
//     cur, err := collection.Find(context.Background(), filter)
//     if err != nil {
//         return nil, err
//     }
//     defer cur.Close(context.Background())
//     var results []bson.M
//     for cur.Next(context.Background()) {
//         var result bson.M
//         err := cur.Decode(&result)
//         if err != nil {
//             return nil, err
//         }
//         results = append(results, result)
//     }
//     return results, nil
// }

func runAggregation(databaseName, collectionName string, pipeline mongo.Pipeline) ([]bson.M, error) {
	cur, err := Client.Database(databaseName).Collection(collectionName).Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []bson.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

var kdatabaseName string = "goblog"

// SetSchemaValidation sets schema validation rules for a collection
//
//	func SetSchemaValidation(collectionName string, validationRules interface{}) error {
//		coll := Client.Database(kdatabaseName).Collection(collectionName)
//		opts := options.Collection().
//			SetValidator(validationRules)
//		if err := coll.Drop(context.Background()); err != nil { // Drop the collection if it exists
//			return err
//		}
//		if err := Client.Database(kdatabaseName).CreateCollection(context.Background(), collectionName, opts); err != nil {
//			return err
//		}
//		return nil
//	}
//
// // SetSchemaValidation sets schema validation rules for a collection
// func SetSchemaValidation(collectionName string, validationRules interface{}) error {
// 	opts := options.CreateCollection()
// 	opts.Validator = validationRules

//		// Create the collection with validation rules
//		err := DB.CreateCollection(context.Background(), collectionName, opts)
//		return err
//	}
//
// Function to set schema validation rules for a collection
func SetSchemaValidation(collectionName string, validationRules interface{}) error {
	opts := options.CreateCollection()
	opts.Validator = validationRules

	// Create the collection with validation rules
	err := DB.CreateCollection(context.Background(), collectionName, opts)
	return err
}

// Function to retrieve all documents from the specified collection
func ViewAll(collectionName string) ([]interface{}, error) {
	collection := DB.Collection(collectionName)
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		var result interface{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// Function to retrieve a single document from the specified collection based on the filter
func View(collectionName string, filter bson.M) (interface{}, error) {
	collection := DB.Collection(collectionName)
	var result interface{}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}
