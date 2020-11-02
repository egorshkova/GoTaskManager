package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
    if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//connection string
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	//fmt.Println(dbUri)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	//DB migration
	//defer db.Close()
	db.Debug().AutoMigrate(&Account{}, &Task{})
}

func GetDB() *gorm.DB {
	return db
}