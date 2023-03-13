package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
	}
	c.Next()
}

func LogoutSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	err := session.Save()
	if err != nil {
		log.Printf("Removing session while logging out failed: %s", err)
		c.Redirect(http.StatusInternalServerError, "/login")
		c.Abort()
	}
	c.Redirect(http.StatusTemporaryRedirect, "/login")
	c.Abort()
}
