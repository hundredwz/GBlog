package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/model"
	"net/http"
)

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		if err != nil {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}
		token := model.DecodeToken(cookie)
		if !token.Verify() {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}
		c.Set("userId", token.UserId)
		c.Next()
	}
}

func Installed() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.Installed {
			c.Redirect(http.StatusFound, "/install")
			c.Abort()
			return
		}
		c.Next()
	}
}
