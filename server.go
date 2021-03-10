package main

import (
	"github.com/gin-gonic/gin"
	"github.com/specter25/gin-microservice/controller"
	"github.com/specter25/gin-microservice/middlewares"
	"github.com/specter25/gin-microservice/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// server := gin.Default()
	server := gin.New()
	//configure the srevre to use these 2 middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
