package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/specter25/gin-microservice/controller"
	"github.com/specter25/gin-microservice/middlewares"
	"github.com/specter25/gin-microservice/repository"
	"github.com/specter25/gin-microservice/service"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	jwtService      service.JWTService         = service.NewJWTService()
	loginService    service.LoginService       = service.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

//to create a log file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	defer videoRepository.CloseDB()

	// setupLogOutput()

	// server := gin.Default()
	server := gin.New()
	//configure the srevre to use these 2 middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/posts", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid",
				})
			}

		})

		apiRoutes.PUT("/posts/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Updated",
				})
			}

		})

		apiRoutes.DELETE("/posts/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Deleted",
				})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	//We can setup this env variable from ebs sonsole
	port := os.Getenv("PORT")
	// Elastic beanstalk forwards request to port 5000
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
