package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *DB = NewDB()

func main() {
	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"answer": "ok",
			"code":   http.StatusOK,
		})
	})
	router.POST("/register", HandleRegister)
	router.POST("/user", HandleGetUser)
	//router.POST("/login", HandleLogin)
	router.Run(":8888") // default, listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
