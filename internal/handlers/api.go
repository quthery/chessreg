package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	c.JSON(200, nil)

}
