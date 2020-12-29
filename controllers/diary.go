package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/QMCHE/diary-server/models"
	"github.com/gin-gonic/gin"
)

// CreateDiary creates diary
func CreateDiary(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameter(s) can't be empty",
		})
		return
	}

	err := models.InsertDiary(title, content)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Create diary failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Diary has been created",
	})
	return
}

// UpdateDiary updates diary
func UpdateDiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	content := c.PostForm("content")

	err := models.UpdateDiary(id, title, content)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Update diary failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Diary has been updated",
	})
	return
}

// DeleteDiary deletes diary by id
func DeleteDiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Print(id)

	err := models.DeleteDiary(id)
	if err != nil {
		log.Print(err)
	}

}
