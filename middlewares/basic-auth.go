package middlewares

import "github.com/gin-gonic/gin"

//this is a authprization middleware
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"ujjwal": "agarwal",
	})
}
