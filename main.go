package main

import (
	"EVMap/authentication"
	"EVMap/config"
	"EVMap/pages"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB
var cfg config.Config

func main() {
	cfg = config.ParseConfig("config.yml")
	//db = database.ConnectDB(cfg.Database.Address, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password,
	//	cfg.Database.DBName)

	// File handling
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Login page
	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		authentication.LoginHandler(w, r, db, cfg)
	})
	// Register page
	http.HandleFunc("/register/", func(w http.ResponseWriter, r *http.Request) {
		authentication.RegisterHandler(w, r, db, cfg)
	})

	// Home page
	http.HandleFunc("/", pages.IndexHandler)

	fmt.Println("Server is listening on " + cfg.Server.Address + ":" + cfg.Server.Port + "...")
	http.ListenAndServe(cfg.Server.Address+":"+cfg.Server.Port, nil)
}
