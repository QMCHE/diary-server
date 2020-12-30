package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/QMCHE/diary-server/models"
	"github.com/QMCHE/diary-server/utils"
	"github.com/gin-gonic/gin"
)

// Login is login controller
func Login(c *gin.Context) {
	userID := c.PostForm("userId")
	password := c.PostForm("password")

	if strings.Trim(userID, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter(s) can't be empty",
		})
		return
	}

	db := utils.DBConnect()

	err := models.IsUserExists(db, userID, password)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
	})
	return
}

// Register is register controller
func Register(c *gin.Context) {
	name := c.PostForm("name")
	userID := c.PostForm("userId")
	password := c.PostForm("password")

	if strings.Trim(name, " ") == "" || strings.Trim(userID, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter(s) can't be empty",
		})
		return
	}

	db := utils.DBConnect()

	if !models.IsUniqueUserID(db, userID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Duplicate userId",
		})
		return
	}

	err := models.InsertUser(db, name, userID, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insert user failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register success",
	})
	return
}
