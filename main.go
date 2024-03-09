package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	sqldb "main/core/db/sql"
	_ "main/docs" // This is required for Swagger to find your documentation
	conUser "main/features/user"
)

// @title User API
// @description API for user management
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	sqldb.Init()
	InitDataBase()
	router := gin.Default()
	router.Static("/docs", "./docs")
	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", conUser.CreateUser)
			user.PUT("/:id", conUser.UpdateUser)
			user.GET("/:id", conUser.GetUserByID)
			user.GET("", conUser.GetUsers)
			user.GET("/search", conUser.SearchUsers)
			user.DELETE("/:id", conUser.DeleteUser)
			user.DELETE("/", conUser.DeleteAllUsers)

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	defer sqldb.DB.Close()
}
