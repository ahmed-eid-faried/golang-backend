package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	redis "main/core/db/redis"
	redisChat2 "main/core/db/redis/chat2"
	sqldb "main/core/db/sql"
	_ "main/docs" // This is required for Swagger to find your documentation
	conAddress "main/features/address"
	conRedis "main/features/addressRedis"

	// chat "main/features/chat"
	addressmongodb "main/features/addressmongodb"
	conUser "main/features/user"
)

// addressmongodb "main/features/addressmongodb"

// 	redisChat "main/core/db/redis/chat"

// @title User API
// @description API for user management
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	sqldb.Init()
	// chat.InitializeRedis()
	InitDataBase()
	router := gin.Default()
	router.Static("/docs", "./docs")
	// router.Static("/core/views/", "./core/views")
	// // Set the templates directory
	// router.LoadHTMLGlob(filepath.Join("./templates", "*.html"))
	// Set the directory to serve static files (HTML, CSS, JS, etc.)
	router.Static("/static", "./static")
	// Define your routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		// Set the status code to 404
		// c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
		c.File("./templates/404.html")
	})
	v1 := router.Group("/api/v1")
	{
		// Grouping routes related to user management
		user := v1.Group("/users")
		{
			// User routes
			user.POST("/", conUser.CreateUser)
			user.PUT("/:id", conUser.UpdateUser)
			user.GET("/:id", conUser.GetUserByID)
			user.GET("", conUser.GetUsers)
			user.GET("/search", conUser.SearchUsers)
			user.DELETE("/:id", conUser.DeleteUser)
			user.DELETE("/", conUser.DeleteAllUsers)
		}
		// Grouping routes related to address management
		address := v1.Group("/address")
		{
			// Address routes
			address.POST("/", conAddress.CreateAddress)
			address.PUT("/:id", conAddress.UpdateAddress)
			address.GET("/:id", conAddress.GetAddressByID)
			address.GET("", conAddress.GetAddresses)
			address.GET("/search", conAddress.SearchAddresses)
			address.DELETE("/:id", conAddress.DeleteAddress)
			address.DELETE("/", conAddress.DeleteAllAddresses)
		}
		// Grouping routes related to address management Mongodb
		addressMongo := v1.Group("/addressMongodb")
		{
			// Address routes
			addressMongo.POST("/", addressmongodb.CreateAddress)
			addressMongo.PUT("/:id", addressmongodb.UpdateAddress)
			addressMongo.GET("/:id", addressmongodb.GetAddressByID)
			addressMongo.GET("", addressmongodb.GetAddresses)
			addressMongo.GET("/search", addressmongodb.SearchAddresses)
			addressMongo.DELETE("/:id", addressmongodb.DeleteAddress)
			addressMongo.DELETE("/", addressmongodb.DeleteAllAddresses)
		}

		redis.Example()
		// Grouping routes related to address management
		redis := v1.Group("/redis")
		{
			// redis routes
			redis.GET("/get", conRedis.GetValue)
			redis.POST("/cache", conRedis.CacheValue)
			redis.DELETE("/remove", conRedis.RemoveValue)
			redis.PUT("/update", conRedis.UpdateValue)
			redis.GET("/getAll", conRedis.GetAllValues)
			redis.DELETE("/removeAll", conRedis.RemoveAllValues)
			redis.GET("/search", conRedis.SearchKeys)
		}

	}

	// Serve index.html
	// v1.GET("/api/v1/chat2/path", redisChat2.HandleChatPath)
	redisChat2.InitChat(router, v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	defer sqldb.DB.Close()
}

// Grouping routes related to chat (realtime) management
// chatapi := v1.Group("/chatapi")
// {

// // Handle chat endpoints
// // Initialize controllers

// chatController := chat.NewChatController()
// chatapi.GET("/:roomId", chatController.GetChatRoom)
// chatapi.POST("/:roomId", chatController.SendMessage)

// // Handle WebSocket connections
// chatapi.GET("/ws", chat.WebSocketHandler)
// chatapi.POST("/:roomId", redisChat.ChatRoute())

// // Serve static files
// chatapi.Static("/static", "./static")
// // Websocket endpoint
// chatapi.GET("/ws", func(c *gin.Context) {
// 	redisChat.HandleWebsocket(c.Writer, c.Request)
// })
// chatapi.GET("/ws", serveHome)
//////////////////////nnnnnnnnnnmmmmmmm//////////////////////////////

// hub := chat.NewHub()
// go hub.Run()

// chatapi.GET("/", serveHome)
// chatapi.GET("/ws", func(c *gin.Context) {
// 	chat.ServeWs(hub, c.Writer, c.Request)
// })

// // // Swagger API documentation routes
// // r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// server := &http.Server{
// 	Addr:           chat.Addr,
// 	Handler:        gin.Default(),
// 	ReadTimeout:    3 * time.Second,
// 	WriteTimeout:   3 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }

// err := server.ListenAndServe()
// if err != nil {
// 	log.Fatal("ListenAndServe: ", err)
// }
// }

// // Handle WebSocket connections
// router.GET("/api/v1/chat/ws", func(c *gin.Context) {
// 	c.String(http.StatusOK, "WebSocket endpoint. Connect using WebSocket client.")
// })

// router.GET("/api/v1/chatapi/ws/gui", serveHome)
// router.GET("/api/v1/chatapi/ws", redisChat.HandleWebsocket)

// Load HTML page
// router.LoadHTMLFiles("./core/db/redis/chat/index.html")
// router.LoadHTMLFiles("./core/db/redis/chat2/index.html")
// // serveHome serves the home page.
// func serveHome(c *gin.Context) {
// 	c.File("./core/db/redis/chat/index.html")
// }
