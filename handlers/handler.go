package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gotickets/db"
	"gotickets/models"
)

type Handler struct {
	datalink *db.DB
}

func NewHandler(datalink *db.DB) *Handler {
	return &Handler{
		datalink: datalink,
	}
}

func (h *Handler) Register(ctx *gin.Context) {
	var payload models.UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr := models.NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	h.datalink.SetUser(usr)
	ctx.JSON(http.StatusOK, usr)
}

func (h *Handler) Login(ctx *gin.Context) {
	var payload models.UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr, err := h.datalink.GetUserByName(payload.LastName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if usr.Password != payload.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(200, gin.H{"auth": true})
}

func (h *Handler) GetUser(ctx *gin.Context) {
	name := ctx.Query("name")
	usr, er := h.datalink.GetUserByName(name)
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   db.ErrUserNotFound,
		})
		return
	}
	log.Printf("%v is here", usr.FirstName)
	ctx.JSON(http.StatusOK, usr)
}
