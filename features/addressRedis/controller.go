package addressRedis

import (
	"net/http"

	// "github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// model "main/core/gin/model"
	redis "main/core/db/redis"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// @Summary Get value from cache
// @Produce json
// @Param key query string true "Key to fetch from cache"
// @Success 200 {string} string "Value from cache"
// @Router /redis/get [get]
// @Tags Redis Cache MEMORY
func GetValue(c *gin.Context) {
	key := c.Query("key")
	val, err := redis.RedisClient.Get(c, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": val})
}

// @Summary Cache a value
// @Accept json
// @Produce json
// @Param key query string true "Key for the value"
// @Param value query string true "Value to be cached"
// @Success 200 {string} string "Value cached successfully"
// @Router /redis/cache [post]
// @Tags Redis Cache MEMORY
func CacheValue(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	err := redis.RedisClient.Set(c, key, value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value cached successfully"})
}

// @Summary Remove a value from cache
// @Produce json
// @Param key query string true "Key to remove from cache"
// @Success 200 {string} string "Value removed successfully"
// @Router /redis/remove [delete]
// @Tags Redis Cache MEMORY
func RemoveValue(c *gin.Context) {
	key := c.Query("key")
	_, err := redis.RedisClient.Del(c, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value removed successfully"})
}

// @Summary Update a value in cache
// @Accept json
// @Produce json
// @Param key query string true "Key for the value"
// @Param value query string true "New value"
// @Success 200 {string} string "Value updated successfully"
// @Router /redis/update [put]
// @Tags Redis Cache MEMORY
func UpdateValue(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	err := redis.RedisClient.Set(c, key, value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value updated successfully"})
}

// GetAllValues retrieves all keys with their values from the Redis cache.
// @Summary Get all keys with their values from cache
// @Produce json
// @Success 200 {object} map[string]string "All keys with their values"
// @Router /redis/getAll [get]
// @Tags Redis Cache MEMORY
func GetAllValues(c *gin.Context) {
	keys, err := redis.RedisClient.Keys(c, "*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make(map[string]string)
	for _, key := range keys {
		// Get the type of the value stored at the key
		keyType, err := redis.RedisClient.Type(c, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Check if the value is a string
		if keyType == "string" {
			// Get the value if it's a string
			value, err := redis.RedisClient.Get(c, key).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			result[key] = value
		} else {
			// Skip this key if the value is not a string
			result[key] = "<non-string value>"
		}
	}
	c.JSON(http.StatusOK, result)
}

// func GetAllValues(c *gin.Context) {
// 	keys, err := redis.RedisClient.Keys(c, "*").Result()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result := make(map[string]string)
// 	for _, key := range keys {
// 		value, err := redis.RedisClient.Get(c, key).Result()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		result[key] = value
// 	}
// 	c.JSON(http.StatusOK, result)
// }

// @Summary Remove all keys from cache
// @Produce json
// @Success 200 {string} string "All keys removed successfully"
// @Router /redis/removeAll [delete]
// @Tags Redis Cache MEMORY
func RemoveAllValues(c *gin.Context) {
	keys, err := redis.RedisClient.Keys(c, "*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = redis.RedisClient.Del(c, keys...).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All keys removed successfully"})
}

// @Summary Search for keys matching a pattern
// @Produce json
// @Param pattern query string true "Pattern to search for (e.g., 'prefix:*' or '*suffix')"
// @Success 200 {object} []string "List of keys matching the pattern"
// @Router /redis/search [get]
// @Tags Redis Cache MEMORY
func SearchKeys(c *gin.Context) {
	pattern := c.Query("pattern")
	keys, err := redis.RedisClient.Keys(c, pattern).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, keys)
}

// // 	model "main/core/gin/model"
// // @Summary Get All Addresses
// // @Description Get all Addresses
// // @Produce json
// // @Success 200 {array} Address
// // @Router /address [get]
// @Tags Redis Cache momery//
// @Tags Address
// func GetAddresses(c *gin.Context) {
// 	var addresses []Address
// 	rows, err := sqldb.DB.Query("SELECT * FROM address")
// 	if err != nil {
// 		log.Println("Error querying address:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying address"})
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var address Address
// 		if err := rows.Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID); err != nil {
// 			log.Println("Error scanning address row:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning address row"})
// 			return
// 		}
// 		addresses = append(addresses, address)
// 	}
// 	c.JSON(http.StatusOK, addresses)
// }

// // @Summary Get a Address by ID
// // @Description Get a Address by ID
// // @Produce json
// // @Param id path int true "Address ID"
// // @Success 200 {object} Address
// // @Router /address/{id} [get]
// @Tags Redis Cache momery//
// @Tags Address
// func GetAddressByID(c *gin.Context) {
// 	// key := c.Param("key")

// 	// // Cache a value
// 	// if err := redis.GetValue(key,  10*time.Second); err != nil {
// 	// 	fmt.Println("Error caching value:", err)
// 	// 	return
// 	// }

// 	id := c.Param("id")
// 	var address Address
// 	err := sqldb.DB.QueryRow("SELECT * FROM address WHERE id = $1", id).Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
// 			return
// 		}
// 		log.Println("Error querying Address:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying Address"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, address)
// }

// // @Summary Create a Address
// // @Description Create a Address
// // @Accept json
// // @Produce json
// // @Param Address body Address true "Address object"
// // @Success 201
// // @Router /address [post]
// @Tags Redis Cache momery//
// @Tags Address
// func CreateAddress(c *gin.Context) {
// 	keyc := c.Param("key")
// 	key := fmt.Sprintf("%s", keyc)
// 	value := c.Param("value")
// 	valuec := fmt.Sprintf("%s", value)

// 	// Cache a value
// 	if err := redis.CacheValue(key, valuec, 10*time.Second); err != nil {
// 		fmt.Println("Error caching value:", err)
// 		return
// 	}

// 	c.Status(http.StatusCreated)
// }

// // @Summary Update a Address
// // @Description Update a Address by ID
// // @Accept json
// // @Produce json
// // @Param id path int true "Address ID"
// // @Param Address body Address true "Address object"
// // @Success 200
// // @Router /address/{id} [put]
// @Tags Redis Cache momery//
// @Tags Address
// func UpdateAddress(c *gin.Context) {
// 	id := c.Param("id")

// 	var address Address
// 	if err := c.ShouldBindJSON(&address); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := sqldb.DB.Exec("UPDATE address SET  street = $1, city = $2, state = $3, postal_code = $4, user_id = $5 WHERE id = $6", address.Street, address.City, address.State, address.PostalCode, address.UserID, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// // @Summary Delete a Address
// // @Description Delete a Address by ID
// // @Param id path int true "Address ID"
// // @Success 200
// // @Router /address/{id} [delete]
// @Tags Redis Cache momery//
// @Tags Address
// func DeleteAddress(c *gin.Context) {
// 	keyc := c.Param("key")
// 	key := fmt.Sprintf("%s", keyc)
// 	// Remove a cached value
// 	if err := redis.RemoveCachedValue(key); err != nil {
// 		fmt.Println("Error removing cached value:", err)
// 		return
// 	}
// 	c.Status(http.StatusOK)
// }

// // @Summary Remove all Address
// // @Description Remove all Addresses from the database
// // @Produce json
// // @Success 204
// // @Router /address [delete]
// @Tags Redis Cache momery//
// @Tags Address
// func DeleteAllAddresses(c *gin.Context) {
// 	_, err := sqldb.DB.Exec("DELETE FROM address")
// 	if err != nil {
// 		log.Println("Error deleting addresses:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting addresses"})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }

// // @Summary Search address
// // @Description Search address
// // @Produce json
// // @Param keyword query string true "Search keyword"
// // @Success 200 {array} Address
// // @Router /address/search [get]
// @Tags Redis Cache momery//
// @Tags Address
// func SearchAddresses(c *gin.Context) {
// 	keyword := c.Query("keyword")
// 	if keyword == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
// 		return
// 	}

// 	// Query the database for addresses matching the keyword
// 	var addresses []Address
// 	rows, err := sqldb.DB.Query("SELECT * FROM addresses WHERE username ILIKE $1 OR email ILIKE $1", "%"+keyword+"%")
// 	if err != nil {
// 		log.Println("Error querying addresses:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying addresses"})
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var address Address
// 		if err := rows.Scan(&address.ID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.UserID); err != nil {
// 			log.Println("Error scanning user row:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user row"})
// 			return
// 		}
// 		addresses = append(addresses, address)
// 	}

// 	c.JSON(http.StatusOK, addresses)
// }
