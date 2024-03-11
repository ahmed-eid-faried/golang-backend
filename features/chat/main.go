package chat

// import (
// 	"log"
// 	"net/http"
// 	"time"
//
//
//
//
//
//
//
//
//
//
//
//

// 	"github.com/gin-gonic/gin"
// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

var Addr = ":3500"

// // @title Chat Application API
// // @version 1.0
// // @description This is a simple chat application API using WebSockets.
// // @host localhost:z
// // @BasePath /
// func main() {
// 	r := gin.Default()

// 	hub := newHub()
// 	go hub.run()

// 	r.GET("/", serveHome)
// 	r.GET("/ws", func(c *gin.Context) {
// 		serveWs(hub, c.Writer, c.Request)
// 	})

// 	// Swagger API documentation routes
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	server := &http.Server{
// 		Addr:           Addr,
// 		Handler:        r,
// 		ReadTimeout:    3 * time.Second,
// 		WriteTimeout:   3 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// // serveHome serves the home page.
// func serveHome(c *gin.Context) {
// 	c.File("home.html")
// }
