package controllers

import (
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

	user := &models.User{
		UserID:   userID,
		Password: password,
	}

	if !user.IsExists(db) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not exists",
		})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Generate token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access-token": accessToken,
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

	user := &models.User{
		Name:     name,
		UserID:   userID,
		Password: password,
	}

	if !user.IsUniqueUserID(db) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Duplicate userId",
		})
		return
	}

	err = user.Create(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insert user failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

// UpdateUser is controller for updating user
func UpdateUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	value, _ := c.Get("claims")
	claims := value.(*utils.Claims)

	db, err := utils.DBConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect database",
		})
		return
	}

	user := &models.User{
		ID: claims.ID,
	}

	err = user.GetUserByID(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find user",
		})
		return
	}

	if strings.Trim(name, " ") != "" {
		user.Name = name
	}
	if strings.Trim(password, " ") == "" {
		user.Password = password
	}

	err = user.Update(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

// DeleteUser is controller for deleting user
func DeleteUser(c *gin.Context) {
	value, _ := c.Get("claims")
	claims := value.(*utils.Claims)

	db, err := utils.DBConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect database",
		})
		return
	}

	user := &models.User{
		ID: claims.ID,
	}

	err = user.Delete(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
	return
}

// // ReissueAccessToken reissue the token
// func ReissueAccessToken(c *gin.Context) {
// 	refreshToken := c.Request.Header.Get("refresh-token")
// 	accessToken := c.Request.Header.Get("access-token")
// 	if strings.Trim(refreshToken, " ") == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Refresh token is required",
// 		})
// 		return
// 	}

// 	if strings.Trim(accessToken, " ") == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Access token is required",
// 		})
// 		return
// 	}

// 	if utils.IsExpired(refreshToken) {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Refresh token is expired",
// 		})
// 		return
// 	}

// 	claims, err := utils.VerifyToken(refreshToken)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Invalid refresh token",
// 		})
// 		return
// 	}

// 	db, err := utils.DBConnect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to connect database",
// 		})
// 		return
// 	}
// 	user := &models.User{
// 		ID: claims.ID,
// 	}

// 	err = user.GetUserByID(db)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Can't find user",
// 		})
// 		return
// 	}

// 	newAccessToken, err := utils.GenerateAccessToken(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to generate token",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"access-token": newAccessToken,
// 	})
// 	return
// }
