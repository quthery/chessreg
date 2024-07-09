package handlers

import (
	"chessreg/internal/database"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	database *database.Storage
}

func NewHandler(database *database.Storage) *Handler {
	return &Handler{database: database}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", h.Index)
	router.POST("/newUser", h.SaveUser)
	router.GET("/DropTable", h.DropTable)
	return router
}
