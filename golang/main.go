package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "world"})
}

// func main() {
// 	router := gin.New()
// 	router.GET("/hello", hello)
// 	router.Run(":9090")
// }

func main() {
	dsn := "u150251147_ahmedmady:8TVkkDgLtbTuZXD@tcp(localhost)/u150251147_ecommerce?charset=utf8mb4"
	// countRowInPage := 9

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.POST("/address", AddressCreate)
	router.GET("/address", AddressRead)
	router.PATCH("/address", AddressUpdate)
	router.DELETE("/address", AddressDelete)

	router.GET("/", hello)

	// Run server
	router.Run(":8080")
}

// Routes
func AddressUpdate(c *gin.Context) {
	// Handle your route logic here
	addressid := c.PostForm("addressid")
	name := c.PostForm("name")
	city := c.PostForm("city")
	street := c.PostForm("street")
	lat := c.PostForm("lat")
	long := c.PostForm("long")

	data := map[string]interface{}{
		"address_name":   name,
		"address_city":   city,
		"address_street": street,
		"address_lat":    lat,
		"address_long":   long,
	}

	updateData(c, "address", data, " address_id =", addressid)
}
func AddressDelete(c *gin.Context) {
	// Handle your route logic here

	addressid := c.PostForm("addressid")
	DeleteData(c, "address", "address_id=", addressid)
}
func AddressRead(c *gin.Context) {
	// Handle your route logic here
	userid := c.PostForm("userid")
	GetAllData(c, "address", "address_userid=", userid)
}
func AddressCreate(c *gin.Context) {
	// Handle your route logic here
	userid := c.PostForm("userid")
	name := c.PostForm("name")
	city := c.PostForm("city")
	street := c.PostForm("street")
	lat := c.PostForm("lat")
	long := c.PostForm("long")
	data := map[string]interface{}{
		"address_userid": userid,
		"address_name":   name,
		"address_city":   city,
		"address_street": street,
		"address_lat":    lat,
		"address_long":   long,
	}

	InsertData(c, "address", data)

}
