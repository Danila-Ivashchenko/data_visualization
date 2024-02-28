package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *coinHandler) CreateLineChart(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no key in request"})
		return
	}

	err := h.service.CreateLineChart(key, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
