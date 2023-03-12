package messages

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func AddMessage(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.AddFlash(message)
	if err := session.Save(); err != nil {
		log.Printf("Error while saving session in \"AddMessage\": %s", err)
	}
}

func GetMessage(c *gin.Context) []interface{} {
	session := sessions.Default(c)
	message := session.Flashes()
	if message != nil {
		err := session.Save()
		if err != nil {
			log.Printf("Error while saving session in \"GetMessage\": %s", err)
		}
	}
	return message
}
