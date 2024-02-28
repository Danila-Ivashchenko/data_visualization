package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *coinHandler) CreateBar(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no key in request"})
		return
	}

	err := h.service.CreateBarChart(key, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
