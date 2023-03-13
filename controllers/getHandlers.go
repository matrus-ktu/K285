package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	user_session := session.Get("user")
	if user_session != nil {
		c.Redirect(http.StatusMovedPermanently, "/account")
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func RegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", nil)
}

func AccountHandler(c *gin.Context) {
	AuthRequired(c)
	c.HTML(http.StatusOK, "under_construction.html", nil)
}
