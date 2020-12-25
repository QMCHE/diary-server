package main

import (
	"github.com/QMCHE/diary-server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.Run()
}
