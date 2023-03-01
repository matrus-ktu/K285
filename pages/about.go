package pages

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("static/about.html"))
	page.Execute(w, nil)
	return
}
