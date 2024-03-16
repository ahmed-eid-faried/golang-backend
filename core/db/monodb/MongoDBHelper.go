package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	Title    string    `bson:"title"`
	Body     string    `bson:"body"`
	Category string    `bson:"category,omitempty"`
	Likes    int       `bson:"likes,omitempty"`
	Tags     []string  `bson:"tags,omitempty"`
	Date     time.Time `bson:"date,omitempty"`
}

type MongoDBHelper struct {
	client *mongo.Client
}

// var dbName string = "goblog"

func NewMongoDBHelper(connectionString string) (*MongoDBHelper, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoDBHelper{client: client}, nil
}

func (h *MongoDBHelper) Close() error {
	return h.client.Disconnect(context.Background())
}

func (h *MongoDBHelper) GetCollection(collection string) *mongo.Collection {
	return h.client.Database(dbName).Collection(collection)
}

func (h *MongoDBHelper) InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return h.GetCollection(collection).InsertOne(context.Background(), document)
	// return DB.Collection(collection).InsertOne(context.Background(), document)

}

func (h *MongoDBHelper) InsertMany(collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	return h.GetCollection(collection).InsertMany(context.Background(), documents)
}

func (h *MongoDBHelper) FindOne(collection string, filter interface{}) *mongo.SingleResult {
	return h.GetCollection(collection).FindOne(context.Background(), filter)
}

func (h *MongoDBHelper) Find(collection string, filter interface{}) (*mongo.Cursor, error) {
	return h.GetCollection(collection).Find(context.Background(), filter)
}

func (h *MongoDBHelper) UpdateOne(collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	return h.GetCollection(collection).UpdateOne(context.Background(), filter, update)
}

func (h *MongoDBHelper) UpdateMany(collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	return h.GetCollection(collection).UpdateMany(context.Background(), filter, update)
}

func (h *MongoDBHelper) DeleteOne(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return h.GetCollection(collection).DeleteOne(context.Background(), filter)
}

func (h *MongoDBHelper) DeleteMany(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return h.GetCollection(collection).DeleteMany(context.Background(), filter)
}

func (h *MongoDBHelper) Aggregate(collection string, pipeline interface{}) (*mongo.Cursor, error) {
	return h.GetCollection(collection).Aggregate(context.Background(), pipeline)
}

// FilterType defines the type of filter
type FilterType int

const (
	Equal FilterType = iota
	NotEqual
	GreaterThan
	GreaterThanOrEqual
	LessThan
	LessThanOrEqual
	In
	And
	Or
	Nor
	Not
	Regex
	Text
	Where
)

// UpdateOperator defines the type of update operator
type UpdateOperator int

const (
	CurrentDate UpdateOperator = iota
	Inc
	Rename
	Set
	Unset
	AddToSet
	Pop
	Pull
	Push
)

// GenerateFilter generates a BSON filter based on the filter type and value
func GenerateFilter(field string, filterType FilterType, value interface{}) bson.M {
	switch filterType {
	case Equal:
		return bson.M{field: value}
	case NotEqual:
		return bson.M{field: bson.M{"$ne": value}}
	case GreaterThan:
		return bson.M{field: bson.M{"$gt": value}}
	case GreaterThanOrEqual:
		return bson.M{field: bson.M{"$gte": value}}
	case LessThan:
		return bson.M{field: bson.M{"$lt": value}}
	case LessThanOrEqual:
		return bson.M{field: bson.M{"$lte": value}}
	case In:
		return bson.M{field: bson.M{"$in": value}}
	case And:
		return bson.M{"$and": value}
	case Or:
		return bson.M{"$or": value}
	case Nor:
		return bson.M{"$nor": value}
	case Not:
		return bson.M{"$not": value}
	case Regex:
		return bson.M{field: bson.M{"$regex": value}}
	case Text:
		return bson.M{"$text": bson.M{"$search": value}}
	case Where:
		return bson.M{"$where": value}
	default:
		return bson.M{}
	}
}

// GenerateUpdate generates a BSON update based on the update operator and value
func GenerateUpdate(operator UpdateOperator, value interface{}) bson.M {
	switch operator {
	case CurrentDate:
		return bson.M{"$currentDate": value}
	case Inc:
		return bson.M{"$inc": value}
	case Rename:
		return bson.M{"$rename": value}
	case Set:
		return bson.M{"$set": value}
	case Unset:
		return bson.M{"$unset": value}
	case AddToSet:
		return bson.M{"$addToSet": value}
	case Pop:
		return bson.M{"$pop": value}
	case Pull:
		return bson.M{"$pull": value}
	case Push:
		return bson.M{"$push": value}
	default:
		return bson.M{}
	}
}

func examplenew() {

	// filter := GenerateFilter("age", GreaterThan, 30)
	// cursor, err := helper.Find("mydb", "mycollection", filter)
	// if err != nil {
	// 	// Handle error
	// }
	// defer cursor.Close(context.Background())
	// // Iterate over cursor and process documents

}
func DefaultValue(TYPE interface{}, DefaultValue interface{}, NULLValue interface{}) interface{} {
	if TYPE == NULLValue {
		TYPE = DefaultValue
	}
	return TYPE
}
