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
	diary.Use(middlewares.VerifyToken())
	{
		diary.POST("/create", controllers.CreateDiary)
	}

	r.Run(":8080")
}
