package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"douyin/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
		}

		if code != 200 {
			c.JSON(400, gin.H{
				"status_code": code,
				"msg":         "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
