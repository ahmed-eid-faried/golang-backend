package addressmongodb

import (
	"context"
	"log"
	"net/http"

	// "github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	core "main/core"
	mongodb "main/core/db/monodb"

	// model "main/core/gin/model"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// 	mongodb "main/core/db/monodb/DB_mongo"

// 	model "main/core/gin/model"

// GetAddresses retrieves all addresses from MongoDB
// @Summary Get All Addresses
// @Description Get all Addresses
// @Produce json
// @Success 200 {array} Address
// @Router /addressMongodb [get]
// @Tags Address Mongodb
func GetAddresses(c *gin.Context) {
	var addresses []Address
	cursor, err := mongodb.DB.Collection("address").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error querying addresses:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying addresses"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &addresses); err != nil {
		log.Println("Error scanning addresses:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

// GetAddressByID retrieves an address by its ID from MongoDB
// @Summary Get a Address by ID
// @Description Get a Address by ID
// @Produce json
// @Param id path string true "Address ID"
// @Success 200 {object} Address
// @Router /addressMongodb/{id} [get]
// @Tags Address Mongodb
func GetAddressByID(c *gin.Context) {
	id := c.Param("id")
	var address Address
	// mongodb.MonHelper.FindOne("address", mongodb.GenerateFilter("id", mongodb.FilterType(0), id))
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}
	core.Vilad(id)

	err = mongodb.DB.Collection("address").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	c.JSON(http.StatusOK, address)
}

// CreateAddress creates a new address in MongoDB
// @Summary Create a Address
// @Description Create a Address
// @Accept json
// @Produce json
// @Param Address body Address true "Address object"
// @Success 201
// @Router /addressMongodb [post]
// @Tags Address Mongodb
func CreateAddress(c *gin.Context) {

	var address Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	core.Vilad(address)
	result, err := mongodb.DB.Collection("address").InsertOne(context.Background(), address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating address"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// UpdateAddress updates an existing address in MongoDB
// @Summary Update a Address
// @Description Update a Address by ID
// @Accept json
// @Produce json
// @Param id path string true "Address ID"
// @Param Address body Address true "Address object"
// @Success 200
// @Router /addressMongodb/{id} [put]
// @Tags Address Mongodb
func UpdateAddress(c *gin.Context) {
	id := c.Param("id")
	var address Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	core.Vilad(address)
	core.Vilad(id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	_, err = mongodb.DB.Collection("address").ReplaceOne(context.Background(), bson.M{"_id": objID}, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating address"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteAddress deletes an address by its ID from MongoDB
// @Summary Delete a Address
// @Description Delete a Address by ID
// @Param id path string true "Address ID"
// @Success 200
// @Router /addressMongodb/{id} [delete]
// @Tags Address Mongodb
func DeleteAddress(c *gin.Context) {
	id := c.Param("id")
	core.Vilad(id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	_, err = mongodb.DB.Collection("address").DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting address"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteAllAddresses deletes all addresses from MongoDB
// @Summary Remove all Address
// @Description Remove all Addresses from the database
// @Produce json
// @Success 204
// @Router /addressMongodb [delete]
// @Tags Address Mongodb
func DeleteAllAddresses(c *gin.Context) {
	_, err := mongodb.DB.Collection("address").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting addresses"})
		return
	}

	c.Status(http.StatusNoContent)
}

// SearchAddresses searches addresses in MongoDB based on a keyword
// @Summary Search address
// @Description Search address
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} Address
// @Router /addressMongodb/search [get]
// @Tags Address Mongodb
func SearchAddresses(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}
	core.Vilad(keyword)

	filter := bson.M{
		"$or": []bson.M{
			{"street": bson.M{"$regex": keyword, "$options": "i"}},
			{"city": bson.M{"$regex": keyword, "$options": "i"}},
			{"state": bson.M{"$regex": keyword, "$options": "i"}},
			{"postal_code": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	var addresses []Address
	cursor, err := mongodb.DB.Collection("address").Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying addresses"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &addresses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}
