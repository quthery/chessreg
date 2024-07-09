package handlers

import (
	"chessreg/internal/attrs"
	"chessreg/internal/schemas"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	c.JSON(200, nil)
}

func (h *Handler) SaveUser(c *gin.Context) {
	var user schemas.UserInsert

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"username": nil,
			"age":      nil,
			"UserId":   nil,
		})
		slog.Error("bind error", attrs.Err(err))
		return
	}

	userId, err := h.database.NewUser(user.Username, user.Age)

	if err != nil {
		slog.Error("database error", attrs.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"username": nil,
			"age":      nil,
			"UserId":   nil,
		})

	} else {
		slog.Info("succesfly added user at", slog.Int("UserID", userId))

		c.JSON(200, gin.H{
			"error":    nil,
			"username": user.Username,
			"age":      user.Age,
			"UserId":   userId,
		})
	}

}

func (h *Handler) DropTable(c *gin.Context) {
	code := h.database.DropTable()
	c.JSON(200, gin.H{
		"droppper": code,
	})
	return
}
