package authentication

import (
	"EVMap/config"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, cfg config.Config) {
	page := template.Must(template.ParseFiles("static/register.html"))
	page.Execute(w, nil)
	return
}
