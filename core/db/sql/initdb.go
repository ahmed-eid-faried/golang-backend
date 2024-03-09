package sqldb

import (
	"database/sql"
	"fmt"
	"log"
	// "net/http"

	// "github.com/bxcodec/faker/v3"
	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/docs" // This is required for Swagger to find your documentation
)

var DB *sql.DB

// Init initializes the database connection and checks for the existence of the required database.
func Init() {
	// Database credentials
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

	// Constructing connection string
	connString := fmt.Sprintf("postgres://%s:%s@%s.frankfurt-postgres.render.com/%s", dbUser, dbPass, dbHost, dbName)

	var err error

	// Open a database connection
	// DB, err = sql.Open("postgres", "postgres://amadytech_user:UBp0KyiORwkldz92RWSNJ0Xaus6Xy86Q@dpg-cnlfl4vsc6pc73cdbbb0-a.frankfurt-postgres.render.com/amadytech")
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

	if !rows.Next() {
		_, err := DB.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			log.Fatal("Error creating database:", err)
			return
		}
	}

	log.Println("Connected to the database")
	defer rows.Close()
}





// var DB *sql.DB
var err error

// // Init initializes the database connection
// func Init() {
// 	// Database connection parameters
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "goblog"

// 	// Open a connection to the database
// 	DB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// Close closes the database connection
func Close() {
	DB.Close()
}