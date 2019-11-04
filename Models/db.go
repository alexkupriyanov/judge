package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
		return
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, dbHost, dbName)
	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db = conn
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db.Debug().AutoMigrate(&EventType{}, &Event{}, &Player{}, &Game{}, &GameType{}, &Team{}, &User{}, &Token{})
	//db.Debug().AutoMigrate(&models.Submission{}, &models.Task{}, &models.PrLang{}, &models.Checker{}, &models.Test{}, &models.Limits{})
}

func GetDB() *gorm.DB {
	return db
}
