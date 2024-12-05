package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gotickets/db"
	"gotickets/handlers"
)

func main() {
	handler := handlers.NewHandler(db.NewDB())
	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status":  http.StatusOK,
		})
	})
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/users", handler.GetUser)
	router.Run(":8888") // default, listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
