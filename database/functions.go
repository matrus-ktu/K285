package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(address string, port string, username string, password string, dbname string) *gorm.DB {
	str := "host=" + address + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Vilnius"
	db, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
