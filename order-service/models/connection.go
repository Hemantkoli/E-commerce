package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	passwd := os.Getenv("PASS")
	port := os.Getenv("PORT")
	dbConnStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, passwd, port, dbName)
	conn, err := gorm.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Order{})
	log.Println("Connected to DB :)")
}

func DBInstance() *gorm.DB {
	return db
}