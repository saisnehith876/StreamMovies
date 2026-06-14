
package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/controllers"
	"github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
}
