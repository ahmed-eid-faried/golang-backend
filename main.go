package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/docs" // This is required for Swagger to find your documentation
)

// User struct represents a user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var DB *sql.DB

// @title User API
// @description API for user management
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	Init()
	InitDataBase()
	router := gin.Default()
	router.Static("/docs", "./docs")
	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("", getUsers)
			user.POST("", createUser)
			user.PUT("/:id", updateUser)
			user.DELETE("/:id", deleteUser)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	defer DB.Close()
}

// @Summary Get all users
// @Description Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
	var users []User
	rows, err := DB.Query("SELECT id, username, email FROM users")
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

	_, err := DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
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

	_, err := DB.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, id)
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

	_, err := DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func Init() {
	dbDriver := "postgres"
	dbUser := "amadytech_user"
	dbPass := "UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q"
	dbHost := "dpg-cnlfl4vsc6pc73cdbbb0-a"
	// dbPort := 5432
	dbName := "amadytech"

	// Internal Database URL
	// postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a/amadytech
	// connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// connString := fmt.Sprintf("postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a/amadytech")

	// postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a:5432/amadytech?sslmode=disable
	// connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	// External Database URL
	// postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a.frankfurt-postgres.render.com/amadytech
	// connString := fmt.Sprintf("postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a.frankfurt-postgres.render.com/amadytech")
	connString := fmt.Sprintf("postgres://%s:%s@%s.frankfurt-postgres.render.com/%s", dbUser, dbPass, dbHost, dbName)
	var err error

	DB, err = sql.Open(dbDriver, connString)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
		return
	}

	rows, err := DB.Query("SELECT datname FROM pg_database WHERE datname = $1", dbName)
	if err != nil {
		log.Fatal("Error checking database existence:", err)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		_, err := DB.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			log.Fatal("Error creating database:", err)
			return
		}
	}

	log.Println("Connected to the database")
}

func InitDataBase() {
	CreateTable("users", "id SERIAL PRIMARY KEY, username VARCHAR(255), email VARCHAR(255)")
}

func CreateTable(tableName string, columns string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, columns)
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
