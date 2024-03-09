package user

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

// @Summary Get all users
// @Description Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []User
	rows, err := sqldb.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error querying users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying users"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Println("Error scanning user row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user row"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	err := sqldb.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Println("Error querying user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create a user
// @Description Create a user
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201
// @Router /users [post]
func CreateUser(c *gin.Context) {
	// // @Param Username path string true "Username"
	// // @Param Email path string true "Email"
	// 	Username := c.Param("Username")
	// 	Email := c.Param("Email")
	// 	_, err := DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", Username, Email)
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error parsing user data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	_, err := sqldb.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		log.Println("Error inserting user data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Update a user
// @Description Update a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User object"
// @Success 200
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := sqldb.DB.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Param id path int true "User ID"
// @Success 200
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := sqldb.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Remove all users
// @Description Remove all users from the database
// @Produce json
// @Success 204
// @Router /users [delete]
func DeleteAllUsers(c *gin.Context) {
	_, err := sqldb.DB.Exec("DELETE FROM users")
	if err != nil {
		log.Println("Error deleting users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting users"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search users
// @Description Search users by username or email
// @Produce json
// @Param keyword query string true "Search keyword (username or email)"
// @Success 200 {array} User
// @Router /users/search [get]
func SearchUsers(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	// Query the database for users matching the keyword
	var users []User
	rows, err := sqldb.DB.Query("SELECT * FROM users WHERE username ILIKE $1 OR email ILIKE $1", "%"+keyword+"%")
	if err != nil {
		log.Println("Error querying users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying users"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Println("Error scanning user row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user row"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }
