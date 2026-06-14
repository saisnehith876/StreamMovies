package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/controllers"
)

func SetupUnProtectedRoutes(router *gin.Engine) {
	//router.Use(middleware.AuthMiddleWare()) -- not needed

	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.GET("/movies", controller.GetMovies())

	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate())
}
