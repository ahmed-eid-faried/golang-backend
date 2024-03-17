package address

import (
	"database/sql"
	"log"
	"net/http"

	// "github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	sqldb "main/core/db/sql"
	// model "main/core/gin/model"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// 	model "main/core/gin/model"

// @Summary Get All Addresses
// @Description Get all Addresses
// @Produce json
// @Success 200 {array} Address
// @Router /address [get]
// @Tags Address
func GetAddresses(c *gin.Context) {
	var addresses []Address
	rows, err := sqldb.DB.Query("SELECT * FROM address")
	if err != nil {
		log.Println("Error querying address:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying address"})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var address Address
		if err := rows.Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID); err != nil {
			log.Println("Error scanning address row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning address row"})
			return
		}
		addresses = append(addresses, address)
	}
	c.JSON(http.StatusOK, addresses)
}

// @Summary Get a Address by ID
// @Description Get a Address by ID
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} Address
// @Router /address/{id} [get]
// @Tags Address
func GetAddressByID(c *gin.Context) {
	id := c.Param("id")
	var address Address
	err := sqldb.DB.QueryRow("SELECT * FROM address WHERE id = $1", id).Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		log.Println("Error querying Address:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying Address"})
		return
	}
	c.JSON(http.StatusOK, address)
}

// @Summary Create a Address
// @Description Create a Address
// @Accept json
// @Produce json
// @Param Address body Address true "Address object"
// @Success 201
// @Router /address [post]
// @Tags Address
func CreateAddress(c *gin.Context) {
	// // @Param Username path string true "Username"
	// // @Param Email path string true "Email"
	// 	Username := c.Param("Username")
	// 	Email := c.Param("Email")
	// 	_, err := DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", Username, Email)
	var address Address
	if err := c.ShouldBindJSON(&address); err != nil {
		log.Println("Error parsing address data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	_, err := sqldb.DB.Exec("INSERT INTO address (id, street,city, state,postal_code, user_id) VALUES ($1, $2, $3, $4, $5, $6)", address.ID, address.Street, address.City, address.State, address.PostalCode, address.UserID)
	if err != nil {
		log.Println("Error inserting Address data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Update a Address
// @Description Update a Address by ID
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Param Address body Address true "Address object"
// @Success 200
// @Router /address/{id} [put]
// @Tags Address
func UpdateAddress(c *gin.Context) {
	id := c.Param("id")

	var address Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := sqldb.DB.Exec("UPDATE address SET  street = $1, city = $2, state = $3, postal_code = $4, user_id = $5 WHERE id = $6", address.Street, address.City, address.State, address.PostalCode, address.UserID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a Address
// @Description Delete a Address by ID
// @Param id path int true "Address ID"
// @Success 200
// @Router /address/{id} [delete]
// @Tags Address
func DeleteAddress(c *gin.Context) {
	id := c.Param("id")
	_, err := sqldb.DB.Exec("DELETE FROM address WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// @Summary Remove all Address
// @Description Remove all Addresses from the database
// @Produce json
// @Success 204
// @Router /address [delete]
// @Tags Address
func DeleteAllAddresses(c *gin.Context) {
	_, err := sqldb.DB.Exec("DELETE FROM address")
	if err != nil {
		log.Println("Error deleting addresses:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting addresses"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search address
// @Description Search address
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} Address
// @Router /address/search [get]
// @Tags Address
func SearchAddresses(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	// Query the database for addresses matching the keyword
	var addresses []Address
	rows, err := sqldb.DB.Query("SELECT * FROM address WHERE street ILIKE $1 OR city ILIKE $1 OR state ILIKE $1 OR postal_code ILIKE $1 OR user_id ILIKE $1 OR id ILIKE $1", "%"+keyword+"%")
	if err != nil {
		log.Println("Error querying address:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying address"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address Address
		if err := rows.Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID); err != nil {
			log.Println("Error scanning user row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user row"})
			return
		}
		addresses = append(addresses, address)
	}

	c.JSON(http.StatusOK, addresses)
}
