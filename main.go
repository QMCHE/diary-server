package main

import (
	"github.com/QMCHE/diary-server/controllers"
	"github.com/QMCHE/diary-server/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	diary := r.Group("/diary")
	{
		diary.POST("/create", middlewares.VerifyToken, controllers.CreateDiary)
	}

	r.Run(":8080")
}
