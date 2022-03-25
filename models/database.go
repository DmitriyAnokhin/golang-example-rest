package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbHost string
var dbPort string
var dbDatabase string
var dbUser string
var dbPassword string
var dsn string

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")

	dsn = "host=" + dbHost + " " +
		"port=" + dbPort + " " +
		"dbname=" + dbDatabase + " " +
		"user=" + dbUser + " " +
		"password=" + dbPassword + " " +
		"sslmode=disable" + " " +
		"TimeZone=UTC"
}

func ConnectDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	db.AutoMigrate(&Track{})

	return db
}
