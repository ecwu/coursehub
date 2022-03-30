package main

import (
	"coursehub/controllers"
	"coursehub/db"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	db.Init()
	router := gin.Default()
	courseV1 := router.Group("api/v1/courses")
	{
		courseV1.GET("/", logger.SetLogger(), controllers.FetchAllCourse)
		courseV1.GET("/:id", logger.SetLogger(), controllers.FetchSingleCourse)
		courseV1.POST("/", logger.SetLogger(), controllers.CreateCourse)
		courseV1.PUT("/:id", logger.SetLogger(), controllers.UpdateCourse)
		courseV1.DELETE("/:id", logger.SetLogger(), controllers.DeleteCourse)
	}
	groupV1 := router.Group("api/v1/groups")
	{
		groupV1.GET("/", logger.SetLogger(), controllers.FetchAllGroup)
		groupV1.POST("/", logger.SetLogger(), controllers.CreateGroup)
	}
	router.Run()

	defer db.CloseDB()
}
