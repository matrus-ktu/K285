package authentication

import (
	"EVMap/config"
	"EVMap/messages"
	"crypto/sha256"
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, cfg config.Config) {
	page := template.Must(template.ParseFiles("static/login.html"))
	if r.Method != http.MethodPost {
		page.Execute(w, nil)
		return
	}

	email := r.FormValue("email")
	password := sha256.Sum256([]byte(r.FormValue("password")))
	passwordString := string(password[:])

	user := Users{}
	db.Where("email = ?", email).First(&user)
	fmt.Println(user)
	if user.Email == "" {
		msg := messages.SysMessage{Success: false, Message: "Nurodytas el. pašto adresas neegzistuoja!"}
		page.Execute(w, msg)
		return
	} else if passwordString != user.Password {
		msg := messages.SysMessage{Success: false, Message: "Slaptažodis neteisingas!"}
		page.Execute(w, msg)
		return
	} else {
		http.Redirect(w, r, fmt.Sprintf("https://%s:%s/%s", cfg.Server.Address, cfg.Server.Port,
			"/about"), http.StatusMovedPermanently)
		return
	}
	page.Execute(w, nil)
	return
}
