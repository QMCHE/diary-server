package controllers

import (
	"net/http"
	"strings"

	"github.com/QMCHE/diary-server/models"
	"github.com/gin-gonic/gin"
)

// Login is login controller
func Login(c *gin.Context) {
	userID := c.PostForm("userId")
	password := c.PostForm("password")

	// Check the parameters is not empty
	if strings.Trim(userID, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameter(s) can't be empty",
		})
		return
	}

	// Check the user not exists
	if !models.IsUserExists(userID, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication failed",
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

	// Check the parameters is not empty
	if strings.Trim(name, " ") == "" || strings.Trim(userID, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parameter(s) can't be empty",
		})
		return
	}

	// Check the userID is unique
	if !models.IsUniqueUserID(userID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Duplicated",
		})
		return
	}

	err := models.InsertUser(name, userID, password)

	// If user insertion fails
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Register failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
	return
}
