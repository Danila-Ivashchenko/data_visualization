package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *coinHandler) UploadFile(c *gin.Context) {
	file, hendler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := h.service.SaveCsvFile(hendler.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ReadCsvData(key, hendler.Filename, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key})
}
