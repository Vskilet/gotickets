package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gotickets/config"
	"gotickets/db"
	"gotickets/handlers"
)

func main() {
	var (
		srvUsers    *http.Server
		srvConcerts *http.Server
	)

	config := config.New()

	go func() {
		users := gin.Default()
		handler := handlers.NewHandler(db.NewDB())
		handler.InitRoutes(users)
		srvUsers = &http.Server{
			Addr:    ":" + strconv.Itoa(config.PortSrvUsers),
			Handler: users.Handler(),
		}
		srvUsers.ListenAndServe()
	}()

	go func() {
		concerts := gin.Default()
		concerts.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
		})
		srvConcerts = &http.Server{
			Addr:    ":" + strconv.Itoa(config.PortSrvConcerts),
			Handler: concerts.Handler(),
		}
		srvConcerts.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := srvConcerts.Shutdown(ctx); err != nil {
		log.Fatal("Concerts server: ", err)
	}
	if err := srvUsers.Shutdown(ctx); err != nil {
		log.Fatal("Users server: ", err)
	}
	log.Println("Server exited")
}
