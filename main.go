package main

import (
	"github.com/gin-gonic/gin"

	"gotickets/db"
	"gotickets/handlers"
)

func main() {
	handler := handlers.NewHandler(db.NewDB())
	router := gin.Default()
	handler.InitRoutes(router)
	router.Run(":8888") // default, listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
