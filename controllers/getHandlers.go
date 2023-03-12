package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func RegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", nil)
}
