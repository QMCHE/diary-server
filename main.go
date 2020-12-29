package main

import (
	"github.com/QMCHE/diary-server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.POST("/create", controllers.CreateDiary)
	r.PUT("/update/:id", controllers.UpdateDiary)
	r.DELETE("/delete/:id", controllers.DeleteDiary)

	r.Run()
}
