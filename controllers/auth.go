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
	user := &models.User{
		UserID:   userID,
		Password: password,
	}

	err := user.IsUserExists(db)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not exists",
		})
		return
	}

	token, err := utils.CreateUserToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Generate token failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
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
	user := &models.User{
		Name:     name,
		UserID:   userID,
		Password: password,
	}

	if !user.IsUniqueUserID(db) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Duplicate userId",
		})
		return
	}

	err := user.CreateUser(db)
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
