package pages

import (
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("static/contact.html"))
	page.Execute(w, nil)
	return
}
