package main

import (
	"EVMap/api"
	"EVMap/config"
	routes "EVMap/controllers"
	"EVMap/database"
	"EVMap/users"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB
var cfg config.Config

func main() {
	// Initialization
	gin.SetMode(gin.ReleaseMode)
	cfg = config.ParseConfig("config.yml")
	db = database.ConnectDB(cfg.Database.Address, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password,
		cfg.Database.DBName)
	err := db.AutoMigrate(&users.Users{}, &users.Organization{})
	if err != nil {
		log.Printf("Database migration failed: %s", err)
		os.Exit(-1)
	}

	// Router configuration
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte(cfg.Server.Secret))))

	// Routes
	router.GET("/", routes.IndexHandler)
	router.GET("/login", routes.LoginHandler)
	router.POST("/login", func(c *gin.Context) {
		routes.LoginHandlerPost(c, db)
	})
	router.GET("/register", routes.RegisterHandler)
	router.POST("/register", func(c *gin.Context) {
		routes.RegisterHandlerPost(c, db)
	})
	router.GET("/account", routes.AccountHandler)
	router.GET("/logout", routes.LogoutSession)
	router.NoRoute(routes.NotFound)

	// API Routes
	router.GET("/api/testroute", api.TestAPI)

	fmt.Println("Server is listening on " + cfg.Server.Address + ":" + cfg.Server.Port + "...")
	err = router.Run(cfg.Server.Address + ":" + cfg.Server.Port)
	if err != nil {
		log.Printf("Server launch has failed: %s", err)
		os.Exit(-1)
	}
}
