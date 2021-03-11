package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/specter25/gin-microservice/service"
)

//this is a authprization middleware
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"ujjwal": "agarwal",
	})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[name]:", claims["name"])
			log.Println("Claims[admin]:", claims["admin"])
			log.Println("Claims[issuer]:", claims["iss"])
			log.Println("Claims[iss]:", claims["iat"])
			log.Println("Claims[exp]:", claims["exp"])

		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
