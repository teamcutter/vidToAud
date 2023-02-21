package middlewares

import "github.com/gin-gonic/gin"

func Authorization(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if token != auth {
			c.AbortWithStatusJSON(401, gin.H{
				"message" : "Invalid authorization token",
			})
			return
		}
		c.Next() // переходим к обработчику
	}
}