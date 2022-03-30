package main

import (
	"coursehub/controllers"
	"coursehub/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()
	v1 := router.Group("api/v1/courses")
	{
		v1.GET("/", controllers.FetchAllCourse)
	}
	router.Run()

	defer db.CloseDB()
}
