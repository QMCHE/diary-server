package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/QMCHE/diary-server/models"
	"github.com/QMCHE/diary-server/utils"
	"github.com/gin-gonic/gin"
)

// GetDiary is getting all diary controller
func GetDiary(c *gin.Context) {
	sort := c.PostForm("sort")
	direction := c.PostForm("direction")
	perPage, _ := strconv.Atoi(c.PostForm("per_page"))
	page, _ := strconv.Atoi(c.PostForm("page"))

	if strings.Trim(sort, " ") == "" {
		sort = "created_at"
	}
	if strings.Trim(direction, " ") == "" {
		direction = "desc"
	}
	if perPage <= 0 {
		perPage = 50
	}
	if page <= 0 {
		page = 1
	}

	db, err := utils.DBConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
	}

	diaries, err := models.GetDiary(db, sort, direction, perPage, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get diaries failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"diaries": diaries,
	})
	return
}

// CreateDiary is creating diary controller
func CreateDiary(c *gin.Context) {
	claims, isExists := c.Get("claims")
	if !isExists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can't get user information",
		})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameter(s) can't be empty",
		})
		return
	}

	db, err := utils.DBConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
	}

	diary := &models.Diary{
		Title:   title,
		Content: content,
		UserID:  claims.(*utils.Claims).ID,
	}

	err = diary.Create(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Diary inserts failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"diary": diary,
	})
	return
}

// UpdateDiary updates diary
func UpdateDiary(c *gin.Context) {
	idString, _ := strconv.Atoi(c.PostForm("id"))
	id := uint(idString)
	// title := c.PostForm("title")
	// content := c.PostForm("content")

	db, err := utils.DBConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect database",
		})
		return
	}

	diary := &models.Diary{
		ID: id,
	}
	// err = diary.GetDiaryByID(db)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to get diary",
	// 	})
	// 	return
	// }
	err = diary.Update(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update diary",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"diary": diary,
	})
	return
}
