package pages

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("static/index.html"))
	page.Execute(w, nil)
	return
}
