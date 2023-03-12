package controllers

import (
	"EVMap/generator"
	"EVMap/messages"
	"EVMap/users"
	"crypto/sha256"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func LoginHandlerPost(c *gin.Context, db *gorm.DB) {
	session := sessions.Default(c)
	message := messages.GetMessage(c)
	email := c.PostForm("email")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(c.PostForm("password"))))

	if message != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"color":   "alert alert-warning",
			"message": message,
		})
	}

	user := users.Users{}
	db.Where("email = ?", email).First(&user)
	if user.Email == "" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"type":    "alert alert-danger",
			"message": "Nurodytas el. pašto adresas neegzistuoja!",
		})
		c.Abort()
	} else if password != user.Password {
		messages.AddMessage(c, "Slaptažodis neteisingas!")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"type":    "alert alert-danger",
			"message": "Slaptažodis neteisingas!",
		})
		c.Abort()
	} else {
		session.Set("user", user.SessionID)
		err := session.Save()
		if err != nil {
			log.Printf("User session saving failed: %s", err)
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"type":    "alert alert-warning",
				"message": "Prisijungti nepavyko. Susisiekite su administratoriumi!",
			})
		}
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	}
}

func RegisterHandlerPost(c *gin.Context, db *gorm.DB) {
	first_name := c.PostForm("firstname")
	last_name := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	repeatPassword := c.PostForm("repeatPassword")

	user := users.Users{}
	db.Where("email = ?", email).First(&user)
	if user.Email != "" {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"type":    "alert alert-warning",
			"message": "El. pašto adresas yra užimtas!",
		})
		c.Abort()
	} else if password != repeatPassword {
		log.Println(password)
		log.Println(repeatPassword)
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"type":    "alert alert-warning",
			"message": "Slaptažodžiai nesutampa!",
		})
		c.Abort()
	} else {
		user.FirstName = first_name
		user.LastName = last_name
		user.Email = email
		user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		user.SessionID = generator.GenerateAlpha(25)
		db.Create(&user)
		db.Commit()
		messages.AddMessage(c, "Paskyra sukurta! Galite prisijungti.")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Next()
	}

}
