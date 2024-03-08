package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/docs" // This is required for Swagger to find your documentation
)

// "main/core/db/sql/mysqldb"
// User struct represents a user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var db *sql.DB

// @title User API
// @description API for user management
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {

	Init()
	router := gin.Default()
	// Serve static files from the "docs" directory
	router.Static("/docs", "./docs")
	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			// Routes
			user.GET("", getUsers)
			user.POST("", createUser)
			user.PUT("/:id", updateUser)
			user.DELETE("/:id", deleteUser)
		}
	}
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Start server
	router.Run(":8080")
	defer db.Close()

}

// @Summary Get all users
// @Description Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {

	// mysqldb.ViewAll("users")
	var users []User

	rows, err :=
		//  mysqldb.ViewAll("users")
		db.Query("SELECT id, username, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Create a user
// @Description Create a user
// @Accept json
// @Produce json
// @Param ID path int true "User ID"
// @Param Username path string true "Username"
// @Param Email path string true "Email"
// @Success 201
// @Router /users [post]
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	if err != nil {
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
func updateUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE users SET username=?, email=? WHERE id=?", user.Username, user.Email, id)
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
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
func Init() {
	// Database connection parameters
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"

	// Open a connection to the database
	connString := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	var err error
	db, err = sql.Open(dbDriver, connString)
	if err != nil {
		fmt.Println("Error opening database connection: %v", err)
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging database: %v", err)
	}

	fmt.Println("Connected to the database")
}
