package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func createDiary(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameter(s) can't be empty",
		})
	}
}
