package controllers

import (
	"EVMap/messages"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		messages.AddMessage(c, "Norėdami tęsti, prisijunkite.")
		c.Redirect(http.StatusUnauthorized, "/login")
		c.Abort()
		return
	}
	c.Next()
}

func LogoutSession() {

}
