package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return

			//it wont continue for eg: if user is not properly authenticated, it wont continue to call tareted ednpoint
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token povided"})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID) //
		c.Set("role", claims.Role)
		c.Next()

	}
}

