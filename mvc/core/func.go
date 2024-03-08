package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
   ctx.JSON(http.StatusOK, gin.H{"message": "world"})
}

func main() {
	router := gin.New()
	router.GET("/hello", hello)
	router.Run(":9090")
}
