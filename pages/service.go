package pages

import (
	"html/template"
	"net/http"
)

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("static/service.html"))
	page.Execute(w, nil)
	return
}
